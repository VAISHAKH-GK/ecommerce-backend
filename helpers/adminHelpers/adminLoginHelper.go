package adminHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var db = databaseConnection.Db
var ctx = databaseConnection.Ctx

func GetAdminUser(userName string, adminUser *models.AdminUser) {
	db.Collection("admin").FindOne(ctx, bson.M{"userName": userName}).Decode(adminUser)
}

func ComparePassword(hashedPassword string, password string) []byte {
	// encoding hashed password and password that use enterd
	var encodedHashedPassowrd = []byte(hashedPassword)
	var encodedPassowrd = []byte(password)
	// comparing password and user entered password
	var err = bcrypt.CompareHashAndPassword(encodedHashedPassowrd, encodedPassowrd)
	if err == nil {
		var details = map[string]interface{}{"status": true}
		var res = helpers.EncodeJson(details)
		return res
	} else {
		var details = map[string]interface{}{"status": false, "reason": "Wrong Password"}
		var res = helpers.EncodeJson(details)
		return res
	}
}
