package helpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/model"
)

var db = model.Db
var ctx = model.Ctx

func DoUserSignUp(body []byte) []byte {
	var data map[string]interface{}
	DecodeJson(body, &data)
	var emailAvailble = CheckIsEmailAvailble(data["email"].(string))
	if !emailAvailble["status"].(bool) {
		var res = EncodeJson(emailAvailble)
		return res
	}
	data["password"] = hashPassword(data["password"])
	var insertedID = insertUser(data)
	var encodedInsertedID = EncodeJson(map[string]interface{}{"status": true, "reason": insertedID})
	return encodedInsertedID
}

func DoUserLogin(body []byte) []byte {
	var data map[string]interface{}
	DecodeJson(body, &data)
	var user map[string]interface{}
	var email string = data["email"].(string)
	getUser(email, &user)
	var status = checkPassword(user["password"].(string), data["password"].(string))
	return status
}
