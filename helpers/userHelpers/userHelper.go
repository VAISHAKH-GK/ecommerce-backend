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
	var user models.User
	// decoding request body
	helpers.DecodeJson(body, &user)
	// check if email is availble
	var emailAvailble = IsEmailAvailble(user.Email)
	if !emailAvailble {
		var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Email in use"})
		return res
	}
	// hashing password
	user.Password = hashPassword(user.Password)
	// inserting user into database
	var insertedID = insertUser(user)
	// response sending to user with user details
	user.Id = insertedID.(primitive.ObjectID)
	var encodedUser = helpers.EncodeJson(map[string]interface{}{"status": true, "user": user})
	return encodedUser
}

// user login
func DoUserLogin(body []byte) []byte {
	var data map[string]interface{}
	// decoding request body
	helpers.DecodeJson(body, &data)
	var user models.User
	// getting user using email
	getUser(data["email"].(string), &user)
	// checking if password is correct
	var status = checkPassword(user.Password, data["password"].(string))
	return status
}
