package controller

import (
	"io"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/adminHelpers"
	"github.com/gorilla/sessions"
)

func AdminLoginRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var body, err = io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res, userId = adminHelpers.DoAdminLogin(body)
	const oneMinute = 60 * 1
	session, err := store.Get(r, "admin")
	helpers.CheckNilErr(err)
	session.Options.MaxAge = oneMinute
	session.Values["userId"] = userId.Hex()
	session.Values["isLoggedIn"] = true
	err = session.Save(r, w)
	helpers.CheckNilErr(err)
	w.Write(res)
}

func AddAdminRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = adminHelpers.CreateAdminUser(body)
	w.Write(res)
}

func AdminCheckLoginRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "admin")
	helpers.CheckNilErr(err)
	var res = adminHelpers.CheckUserLogin(session)
	w.Write(res)
}

func AdminGetUserDataRotue(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	session, err := store.Get(r, "admin")
	helpers.CheckNilErr(err)
	var res = adminHelpers.GetAdminUserData(session)
	w.Write(res)
}
