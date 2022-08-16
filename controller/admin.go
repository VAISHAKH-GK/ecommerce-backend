package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/adminHelpers"
	"github.com/gorilla/sessions"
)

func AdminLoginRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res, userId = adminHelpers.DoAdminLogin(body)
	const oneMinute = 60 * 1
	session, err := store.Get(r, "admin")
	helpers.CheckNilErr(err)
	session.Options.MaxAge = oneMinute
	session.Values["adminUserId"] = userId.Hex()
	session.Values["isAdminLoggedIn"] = true
	err = session.Save(r, w)
	helpers.CheckNilErr(err)
	w.Write(res)
}

func AddAdminRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = adminHelpers.CreateAdminUser(body)
	w.Write(res)
}
