package userHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrders(userId primitive.ObjectID) []map[string]interface{} {
  var cursor,err = db.Collection("order").Find(ctx,bson.M{"userId":userId})
  helpers.CheckNilErr(err)
  var orders []map[string]interface{}
  for cursor.Next(ctx) {
    var order map[string]interface{}
    cursor.Decode(&order)
    orders = append(orders, order)
  }
  return orders
}
