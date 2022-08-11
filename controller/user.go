package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/userHelpers"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// post request on /api/user/signup
func UserSignUpRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = userHelpers.DoUserSignUp(body)
	w.Write(res)
}

// post request on /api/user/login
func UserLoginRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res, userId = userHelpers.DoUserLogin(body)
	session, err := store.Get(r, "user")
	helpers.CheckNilErr(err)
	session.Values["userId"] = userId.Hex()
	err = session.Save(r, w)
	helpers.CheckNilErr(err)
	w.Write(res)
}

// request on /api/user/checklogin
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	session, err := store.Get(r, "user")
	helpers.CheckNilErr(err)
	userId, err := primitive.ObjectIDFromHex(session.Values["userId"].(string))
	fmt.Println(session.Values)
	if userId != primitive.NilObjectID {
		fmt.Println(userId)
		var res = helpers.EncodeJson(map[string]interface{}{"status": true})
		w.Write(res)
	} else {
		var res = helpers.EncodeJson(map[string]interface{}{"status": false})
		w.Write(res)
	}
}
