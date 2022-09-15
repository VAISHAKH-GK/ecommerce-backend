package userHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var w sync.WaitGroup

func getCartProducts(userId primitive.ObjectID) []map[string]interface{} {
	var cursor = db.Collection("cart").FindOne(ctx, bson.M{"userId": userId}, options.FindOne().SetProjection(bson.M{"products": 1, "_id": 0}))
	var cart map[string][]map[string]interface{}
	var err = cursor.Decode(&cart)
	helpers.CheckNilErr(err)
	return cart["products"]
}

func createOrderDetails(order map[string]interface{}, userId primitive.ObjectID, products []map[string]interface{}, total int) map[string]interface{} {
	var paymentMethod = order["payMethod"].(string)
	var status string
	if paymentMethod == "COD" {
		status = "placed"
	} else {
		status = "pending"
	}
	var orderDetails = map[string]interface{}{
		"deliveryDetails": map[string]interface{}{
			"address": order["address"].(string),
			"mobile":  order["number"].(string),
			"email":   order["email"].(string),
			"name":    order["name"].(string),
		},
		"userId":        userId,
		"paymentMethod": paymentMethod,
		"products":      products,
		"status":        status,
		"total":         total,
		"date":          order["date"],
	}
	return orderDetails
}

func addOrder(orderDetails map[string]interface{}) {
	w.Add(2)
	go func() {
		var _, err = db.Collection("order").InsertOne(ctx, orderDetails)
		helpers.CheckNilErr(err)
		w.Done()
	}()
	go func() {
		var _, err = db.Collection("cart").DeleteOne(ctx, bson.M{"userId": orderDetails["userId"].(primitive.ObjectID)})
		helpers.CheckNilErr(err)
		w.Done()
	}()
	w.Wait()
}
