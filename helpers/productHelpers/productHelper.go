package productHelpers

import (
	"fmt"

	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
)

var db = databaseConnection.Db
var ctx = databaseConnection.Ctx

func AddNewProduct(body []byte) []byte {
	var product models.Product
	helpers.DecodeJson(body, &product)
	fmt.Println("this is product")
	fmt.Println(product)
	var id = addProduct(product)
	var response = helpers.EncodeJson(map[string]interface{}{"status": true, "id": id})
	return response
}
