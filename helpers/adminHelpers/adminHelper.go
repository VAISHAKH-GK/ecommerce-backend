package adminHelpers

import (
	"encoding/json"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DoAdminLogin(body []byte) ([]byte, primitive.ObjectID) {
	var data map[string]interface{}
	// decoding request body
	helpers.DecodeJson(body, &data)
	var user models.AdminUser
	// getting user using email
	GetAdminUser(data["userName"].(string), &user)
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

func CreateAdminUser(body []byte) []byte {
	var data map[string]interface{}
	helpers.DecodeJson(body, &data)
	var hashedPassowrd = hashPassword(data["password"])
	data["password"] = hashedPassowrd
	var insertedId = insertUser(data)
	var response = map[string]interface{}{"status": true, "id": insertedId}
	var res, err = json.Marshal(response)
	helpers.CheckNilErr(err)
	return res
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
