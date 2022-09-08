package helpers

import (
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckLogin(session *sessions.Session) bool {
	if session.Values["isLoggedIn"] != true {
		return false
	} else {
		return true
	}
}

func GetUserId(session *sessions.Session) primitive.ObjectID {
	var userId, err = primitive.ObjectIDFromHex(session.Values["userId"].(string))
  CheckNilErr(err)
  return userId
}

func NotLoggedInResponse() []byte {
	var res = EncodeJson(map[string]interface{}{"status": false, "reason": "Not Logged In"})
	return res
}
