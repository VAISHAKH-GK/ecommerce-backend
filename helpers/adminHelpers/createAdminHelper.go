package adminHelpers

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pass interface{}) string {
	var password = pass.(string)
	// generating hash
	var hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helpers.CheckNilErr(err)
	// converting hash to string
	var hashedPassword = string(hash)
	return hashedPassword
}
func insertUser(user any) interface{} {
	var response, err = db.Collection("admin").InsertOne(ctx, user)
	helpers.CheckNilErr(err)
	var id = response.InsertedID
	return id
}
