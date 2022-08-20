package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addProduct(product models.Product) primitive.ObjectID {
	var response, err = db.Collection("product").InsertOne(ctx, product)
	helpers.CheckNilErr(err)
	return response.InsertedID.(primitive.ObjectID)
}
