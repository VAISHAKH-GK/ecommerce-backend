package userHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = databaseConnection.Db
var ctx = databaseConnection.Ctx

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
