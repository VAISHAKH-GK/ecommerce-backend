package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getProdcutsFromCart(userId primitive.ObjectID) primitive.A {
	var cursor, err = db.Collection("user").Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{"_id": userId},
		},
		{
			"$lookup": bson.M{
				"from":         "product",
				"localField":   "cart",
				"foreignField": "_id",
				"as":           "cartItems",
			},
		},
		{
			"$project": bson.M{"cartItems": 1, "_id": 0},
		}})
	helpers.CheckNilErr(err)
	var products primitive.A
	for cursor.Next(ctx) {
		var product map[string]interface{}
		var err = cursor.Decode(&product)
		helpers.CheckNilErr(err)
		products = product["cartItems"].(primitive.A)
	}
	return products
}
