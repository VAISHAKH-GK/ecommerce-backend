package adminHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var db = databaseConnection.Db
var ctx = databaseConnection.Ctx

func GetAdminUser(userName string, adminUser *models.AdminUser) {
	db.Collection("admin").FindOne(ctx, bson.M{"userName": userName}).Decode(adminUser)
}

// check if password matches
func checkPassword(hashedPassword string, password string) bool {
	// encoding hashed password and password that use enterd
	var encodedHashedPassowrd = []byte(hashedPassword)
	var encodedPassowrd = []byte(password)
	// comparing password and user entered password
	var err = bcrypt.CompareHashAndPassword(encodedHashedPassowrd, encodedPassowrd)
	if err == nil {
		return true
	} else {
		return false
	}
}
