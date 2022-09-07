package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addToCart(userId primitive.ObjectID, newProduct map[string]interface{}) {
	var cart map[string]interface{}
	db.Collection("cart").FindOne(ctx, bson.M{"userId": userId}).Decode(&cart)
	if len(cart) != 0 {
		for _, product := range cart["products"].(primitive.A) {
			if product.(map[string]interface{})["productId"].(primitive.ObjectID) == newProduct["productId"].(primitive.ObjectID) {
				var _, err = db.Collection("cart").UpdateOne(ctx, bson.M{"products.productId": newProduct["productId"].(primitive.ObjectID), "userId": userId}, bson.M{"$set": bson.M{"products.$.quantity": newProduct["quantity"].(int)}})
				helpers.CheckNilErr(err)
				return
			}
		}
		var _, err = db.Collection("cart").UpdateOne(ctx, bson.M{"userId": userId}, bson.M{"$push": bson.M{"products": newProduct}})
		helpers.CheckNilErr(err)
		return
	}
	var cartUser = map[string]interface{}{"userId": userId, "products": []interface{}{newProduct}}
	var _, err = db.Collection("cart").InsertOne(ctx, cartUser)
	helpers.CheckNilErr(err)
	return
}
