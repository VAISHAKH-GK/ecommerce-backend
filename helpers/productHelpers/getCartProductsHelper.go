package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getProdcutsFromCart(userId primitive.ObjectID) []map[string]interface{} {
	var cursor, err = db.Collection("cart").Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{"userId": userId},
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
		var err = cursor.Decode(&product)
		helpers.CheckNilErr(err)
		products = append(products, product)
	}
	return products
}

func GetTotalCartAmount(userId primitive.ObjectID) int {
	var cursor, err = db.Collection("cart").Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{"userId": userId},
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
		},
		{
			"$group": bson.M{"_id": nil, "total": bson.M{"$sum": bson.M{"$multiply": []string{"$quantity", "$product.price"}}}},
		}})
	helpers.CheckNilErr(err)
	var totalPrice int
	for cursor.Next(ctx) {
		var response map[string]interface{}
		var err = cursor.Decode(&response)
		helpers.CheckNilErr(err)
		totalPrice = int(response["total"].(int32))
	}
	return totalPrice
}
