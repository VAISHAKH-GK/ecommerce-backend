package helpers

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func CheckIsEmailAvailble(email string) map[string]interface{} {
	var user map[string]interface{}
	db.Collection("user").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if user != nil {
		var res = map[string]interface{}{"status": false, "reason": "email in use"}
		return res
	} else {
		var res = map[string]interface{}{"status": true}
		return res
	}
}

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
