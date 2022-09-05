package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = databaseConnection.Db
var ctx = databaseConnection.Ctx

func AddNewProduct(product models.Product) []byte {
	var id = addProduct(product)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true, "id": id})
	return res
}

func EditProduct(product models.Product, id string) []byte {
	var objectId, err = primitive.ObjectIDFromHex(id)
	helpers.CheckNilErr(err)
	var status = updateProduct(product, objectId)
	var res = helpers.EncodeJson(map[string]interface{}{"status": status})
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

func DeleteProduct(id string) []byte {
	var objectId, err = primitive.ObjectIDFromHex(id)
	helpers.CheckNilErr(err)
	if _, err = db.Collection("product").DeleteOne(ctx, bson.M{"_id": objectId}); err != nil {
		panic(err)
	}
	var res = helpers.EncodeJson(map[string]interface{}{"status": true})
	return res
}

func AddProductToCart(productId primitive.ObjectID, userId primitive.ObjectID) []byte {
	var _, err = db.Collection("user").UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$push": bson.M{"cart": productId}})
	helpers.CheckNilErr(err)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true})
	return res
}

func GetCartProducts(userId primitive.ObjectID) []byte {
	var products = getProdcutsFromCart(userId)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true, "products": products})
	return res
}
