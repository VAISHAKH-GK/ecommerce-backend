package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/userHelpers"
)

var db = databaseConnection.Db
var ctx = databaseConnection.Ctx

// post request on /api/user/signup
func UserSignUpRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = userHelpers.DoUserSignUp(body)
	w.Write(res)
}

// post request on /api/user/login
func UserLoginRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = userHelpers.DoUserLogin(body)
	w.Write(res)
}
