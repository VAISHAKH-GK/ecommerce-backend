package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func getProducts(numberOfProducts int) []map[string]interface{} {
	var cursor, err = db.Collection("product").Find(ctx, bson.M{})
	helpers.CheckNilErr(err)
	var products []map[string]interface{}
	for cursor.Next(ctx) {
		if len(products) < numberOfProducts {
			var product map[string]interface{}
			err = cursor.Decode(&product)
			helpers.CheckNilErr(err)
			products = append(products, product)
		} else {
			break
		}
	}
	return products
}
