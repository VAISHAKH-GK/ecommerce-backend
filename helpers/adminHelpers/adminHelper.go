package adminHelpers

import (
	"encoding/json"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
)

func DoAdminLogin(body []byte) []byte {
	var data map[string]interface{}
	helpers.DecodeJson(body, &data)
  var adminUser models.AdminUser
	GetAdminUser(data["userName"].(string),&adminUser)
	var res = ComparePassword(adminUser.Password, data["password"].(string))
	return res
}

func CreateAdminUser(body []byte) []byte {
  var data map[string]interface{}
  helpers.DecodeJson(body,&data)
  var hashedPassowrd = hashPassword(data["password"])
  data["password"] = hashedPassowrd
  var insertedId = insertUser(data)
  var response = map[string]interface{}{"status":true,"id":insertedId}
  var res,err = json.Marshal(response)
  helpers.CheckNilErr(err)
  return res
}
