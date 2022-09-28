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

func AddProductToCart(userId primitive.ObjectID, productId primitive.ObjectID, count int) []byte {
	addToCart(userId, productId, count)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true})
	return res
}

func GetCartProducts(userId primitive.ObjectID) []byte {
	var products = getProdcutsFromCart(userId)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true, "products": products})
	return res
}

func RemoveCartProduct(userId primitive.ObjectID, productId primitive.ObjectID) []byte {
	db.Collection("cart").UpdateOne(ctx, bson.M{"userId": userId}, bson.M{"$pull": bson.M{"products": bson.M{"productId": productId}}})
	var res = helpers.EncodeJson(map[string]interface{}{"status": true})
	return res
}

func GetTotalAmount(userId primitive.ObjectID) []byte {
	var total = GetTotalCartAmount(userId)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true, "total": total})
	return res
}

func SearchProducts(searchWord string) []models.Product {
	var cursor, err = db.Collection("product").Find(ctx, bson.M{"$text": bson.M{"$search": searchWord}})
	helpers.CheckNilErr(err)
	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	return products
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
