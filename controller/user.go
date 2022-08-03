package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/model"
)

var db = model.Db
var ctx = model.Ctx

// post request on /api/user/signup
func UserSignUpRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = helpers.DoUserSignUp(body)
	w.Write(res)
}

// post request on /api/user/login
func UserLoginRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = helpers.DoUserLogin(body)
	w.Write(res)
}
