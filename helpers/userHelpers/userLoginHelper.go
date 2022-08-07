package userHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// get user from databse using email
func getUser(email string, user *models.User) {
	db.Collection("user").FindOne(ctx, bson.M{"email": email}).Decode(user)
}

//  check if password matches
func checkPassword(hashedPassword string, password string) []byte {
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
