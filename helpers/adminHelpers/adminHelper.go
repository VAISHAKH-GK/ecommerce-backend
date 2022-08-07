package adminHelpers

import (
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
