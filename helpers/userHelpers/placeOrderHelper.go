package userHelpers

import (
	"os"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	razorpay "github.com/razorpay/razorpay-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCartProducts(userId primitive.ObjectID) []map[string]interface{} {
	var cursor = db.Collection("cart").FindOne(ctx, bson.M{"userId": userId}, options.FindOne().SetProjection(bson.M{"products": 1, "_id": 0}))
	var cart map[string][]map[string]interface{}
	var err = cursor.Decode(&cart)
	helpers.CheckNilErr(err)
	return cart["products"]
}

func createOrderDetails(order map[string]interface{}, userId primitive.ObjectID, products []map[string]interface{}, total int) map[string]interface{} {
	var paymentMethod = order["paymentMethod"].(string)
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

func adddOrder(orderDetails map[string]interface{}) primitive.ObjectID {
	var response, err = db.Collection("order").InsertOne(ctx, orderDetails)
	helpers.CheckNilErr(err)
	_, err = db.Collection("cart").DeleteOne(ctx, bson.M{"userId": orderDetails["userId"].(primitive.ObjectID)})
	helpers.CheckNilErr(err)
	return response.InsertedID.(primitive.ObjectID)
}

func createOnlineOrder(orderDetails map[string]interface{}, orderId primitive.ObjectID) (map[string]interface{}, string) {
	var razorpayKey = os.Getenv("RP_KEY")
	var razorpaySecret = os.Getenv("RP_SECRET")
	var client = razorpay.NewClient(razorpayKey, razorpaySecret)
	var data = map[string]interface{}{
		"amount":   orderDetails["total"].(int) * 100,
		"currency": "INR",
		"receipt":  orderId.Hex(),
	}
	var body, err = client.Order.Create(data, nil)
	helpers.CheckNilErr(err)
	return body, razorpayKey
}
