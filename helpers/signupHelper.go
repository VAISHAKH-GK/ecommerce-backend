package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pass interface{}) string {
	var password = pass.(string)
	var hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	CheckNilErr(err)
	var hashedPassword = string(hash)
	return hashedPassword
}

func insertUser(user any) interface{} {
	var response, err = db.Collection("user").InsertOne(ctx, user)
	CheckNilErr(err)
	var id = response.InsertedID
	return id
}
