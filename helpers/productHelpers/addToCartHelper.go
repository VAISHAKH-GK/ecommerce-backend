package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addToCart(userId primitive.ObjectID, productId primitive.ObjectID, count int) {
	var newProduct = map[string]interface{}{"productId": productId, "quantity": count}
	var cart map[string]interface{}
	db.Collection("cart").FindOne(ctx, bson.M{"userId": userId}).Decode(&cart)
	if len(cart) != 0 {
		for _, product := range cart["products"].(primitive.A) {
			if product.(map[string]interface{})["productId"].(primitive.ObjectID) == productId {
				var _, err = db.Collection("cart").UpdateOne(ctx, bson.M{"products.productId": productId, "userId": userId}, bson.M{"$inc": bson.M{"products.$.quantity": count}})
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
