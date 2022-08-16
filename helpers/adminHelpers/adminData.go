package adminHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getUserById(userId primitive.ObjectID, adminUser *models.AdminUser) {
	db.Collection("admin").FindOne(ctx, map[string]interface{}{"_id": primitive.ObjectID(userId)}).Decode(adminUser)
}
