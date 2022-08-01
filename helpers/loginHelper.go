package helpers

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func getUser(email string, user *map[string]interface{}) {
	db.Collection("user").FindOne(ctx, bson.M{"email": email}).Decode(&user)
}

func checkPassword(hashedPassword string, password string) []byte {
	var encodedHashedPassowrd = []byte(hashedPassword)
	var encodedPassowrd = []byte(password)
	var err = bcrypt.CompareHashAndPassword(encodedHashedPassowrd, encodedPassowrd)
	if err == nil {
		var details = map[string]interface{}{"status": true}
		var res = EncodeJson(details)
		return res
	} else {
		var details = map[string]interface{}{"status": false, "reason": "Wrong Password"}
		var res = EncodeJson(details)
		return res
	}
}
