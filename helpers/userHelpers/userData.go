package userHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func getUserById(userId primitive.ObjectID,user *models.User) {
	db.Collection("user").FindOne(ctx, map[string]interface{}{"_id": primitive.ObjectID(userId)}).Decode(user)
}
