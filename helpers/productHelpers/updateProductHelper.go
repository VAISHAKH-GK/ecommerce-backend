package productHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateProduct(product models.Product, id primitive.ObjectID) bool {
	var _,err = db.Collection("product").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": product})
  helpers.CheckNilErr(err)
  return true
}
