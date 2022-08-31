package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = databaseConnection.Db
var ctx = databaseConnection.Ctx

func AddNewProduct(product models.Product) []byte {
	var id = addProduct(product)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true, "id": id})
	return res
}

func GetAllProducts(numberOfProducts int) []byte {
	var products = getProducts(numberOfProducts)
	var res = helpers.EncodeJson(products)
	return res
}

func GetOneProduct(id string) []byte {
	var objectId, err = primitive.ObjectIDFromHex(id)
	helpers.CheckNilErr(err)
	var product = getProduct(objectId)
	var res = helpers.EncodeJson(product)
	return res
}