package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getProduct(id primitive.ObjectID) models.Product {
	var product models.Product
	db.Collection("product").FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	return product
}
