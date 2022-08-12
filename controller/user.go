package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/userHelpers"
	"github.com/gorilla/sessions"
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
	session.Values["isLoggedIn"] = true
	err = session.Save(r, w)
	helpers.CheckNilErr(err)
	w.Write(res)
}

// request on /api/user/checklogin
func CheckLoginRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	session, err := store.Get(r, "user")
	helpers.CheckNilErr(err)
	var res = userHelpers.CheckUserLogin(session)
	w.Write(res)
}

func GetUserDataRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	session, err := store.Get(r, "user")
	helpers.CheckNilErr(err)
	var res = userHelpers.GetUserData(session)
	w.Write(res)
}
