package userHelpers

import (
	"sync"

	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/productHelpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = databaseConnection.Db
var ctx = databaseConnection.Ctx

var waitGroup sync.WaitGroup

// user singup
func DoUserSignUp(body []byte) []byte {
	var user map[string]interface{}
	// decoding request body
	helpers.DecodeJson(body, &user)
	// check if email is availble
	var emailAvailble = IsEmailAvailble(user["email"].(string))
	if !emailAvailble {
		var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Email in use"})
		return res
	}
	// hashing password
	user["password"] = hashPassword(user["password"].(string))
	// inserting user into database
	var insertedID = insertUser(user)
	// response sending to user with user details
	user["id"] = insertedID.(primitive.ObjectID)
	var encodedUser = helpers.EncodeJson(map[string]interface{}{"status": true})
	return encodedUser
}

// user login
func DoUserLogin(body []byte) ([]byte, primitive.ObjectID) {
	var data map[string]interface{}
	// decoding request body
	helpers.DecodeJson(body, &data)
	var user models.User
	// getting user using email
	getUser(data["email"].(string), &user)
	// checking if password is correct
	var status = checkPassword(user.Password, data["password"].(string))
	// sending response login failed
	if !status {
		var res = helpers.EncodeJson(map[string]interface{}{"status": status, "reason": "Login Falied"})
		var userId primitive.ObjectID
		return res, userId
	}
	// sending response if login successfull
	var res = helpers.EncodeJson(map[string]interface{}{"status": status})
	var userId primitive.ObjectID = user.Id
	return res, userId
}

// checking if user already logedIn
func CheckUserLogin(session *sessions.Session) []byte {
	// getting userId from session
	var isLoggedIn = session.Values["isLoggedIn"]
	// checking if userId is nil
	if isLoggedIn == true {
		// creating response and returning
		var res = helpers.EncodeJson(map[string]interface{}{"status": true})
		return res
	} else {
		// creating response and returning
		var res = helpers.EncodeJson(map[string]interface{}{"status": false})
		return res
	}
}

func GetUserData(session *sessions.Session) []byte {
	var id = session.Values["userId"]
	if id == nil {
		var res = helpers.EncodeJson(false)
		return res
	}
	var userId, err = primitive.ObjectIDFromHex(session.Values["userId"].(string))
	helpers.CheckNilErr(err)
	var user models.User
	getUserById(userId, &user)
	var res = helpers.EncodeJson(user)
	return res
}

func CheckLogin(session *sessions.Session) bool {
	if session.Values["isLoggedIn"] != true {
		return false
	} else {
		return true
	}
}

func GetUserId(session *sessions.Session) primitive.ObjectID {
	var userId, err = primitive.ObjectIDFromHex(session.Values["userId"].(string))
	helpers.CheckNilErr(err)
	return userId
}

func NotLoggedInResponse() []byte {
	var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Not Logged In"})
	return res
}

func PlaceOrder(order map[string]interface{}, userId primitive.ObjectID) []byte {
	var products []map[string]interface{}
	var total int
	waitGroup.Add(2)
	go func(products *[]map[string]interface{}) {
		*products = getCartProducts(userId)
		waitGroup.Done()
	}(&products)
	go func(total *int) {
		*total = productHelpers.GetTotalCartAmount(userId)
		waitGroup.Done()
	}(&total)
	waitGroup.Wait()
	var orderDetails = createOrderDetails(order, userId, products, total)
	var orderId = adddOrder(orderDetails)
	if order["paymentMethod"].(string) == "ONLINE" {
		var orderData, key = createOnlineOrder(orderDetails, orderId)
		var res = helpers.EncodeJson(map[string]interface{}{"status": "PENDING", "orderData": orderData, "key": key})
		return res
	} else {
		var res = helpers.EncodeJson(map[string]interface{}{"status": "DONE"})
		return res
	}
}

func GetOrderProducts(orderId primitive.ObjectID) []map[string]interface{} {
	var cursor, err = db.Collection("order").Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{"_id": orderId},
		},
		{
			"$unwind": "$products",
		},
		{
			"$project": bson.M{
				"productId": "$products.productId",
				"quantity":  "$products.quantity",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "product",
				"localField":   "productId",
				"foreignField": "_id",
				"as":           "product",
			},
		},
		{
			"$unwind": "$product",
		},
		{
			"$project": bson.M{"product": 1, "quantity": 1, "_id": 0},
		}})
	helpers.CheckNilErr(err)
	var products []map[string]interface{}
	for cursor.Next(ctx) {
		var product map[string]interface{}
		cursor.Decode(&product)
		products = append(products, product)
	}
	return products
}
