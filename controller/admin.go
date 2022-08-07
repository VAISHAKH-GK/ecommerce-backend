package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/adminHelpers"
)

func AdminLoginRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = adminHelpers.DoAdminLogin(body)
	w.Write(res)
}

func AddAdminUser(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = adminHelpers.CreateAdminUser(body)
	w.Write(res)
}
