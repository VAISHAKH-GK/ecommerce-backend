package userHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// checking is email is availble
func IsEmailAvailble(email string) bool {
	var user map[string]interface{}
	// finding user with email if exists
	db.Collection("user").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if user != nil {
		return false
	} else {
		return true
	}
}

// hashing password
func hashPassword(pass interface{}) string {
	var password = pass.(string)
	// generating hash
	var hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helpers.CheckNilErr(err)
	// converting hash to string
	var hashedPassword = string(hash)
	return hashedPassword
}

// inserting user into database
func insertUser(user any) interface{} {
	var response, err = db.Collection("user").InsertOne(ctx, user)
	helpers.CheckNilErr(err)
	var id = response.InsertedID
	return id
}
