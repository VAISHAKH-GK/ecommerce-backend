package controller

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/adminHelpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/productHelpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/models"
	"github.com/gorilla/sessions"
)

func AdminLoginRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var body, err = io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res, userId = adminHelpers.DoAdminLogin(body)
	const twentyMinute = 60 * 1 * 20
	session, err := store.Get(r, "admin")
	helpers.CheckNilErr(err)
	session.Options.MaxAge = twentyMinute
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

func AdminLogoutRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var sesssion, err = store.Get(r, "admin")
	helpers.CheckNilErr(err)
	sesssion.Options.MaxAge = -1
	sesssion.Save(r, w)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true})
	w.Write(res)
}

func AddProductRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var data models.Product
	helpers.DecodeJson(body, &data)
	var res = productHelpers.AddNewProduct(data)
	w.Write(res)
}

func AddProductImageRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = io.ReadAll(r.Body)
	var id = r.URL.Query().Get("id")
	helpers.CheckNilErr(err)
	ioutil.WriteFile("public/images/"+id+".jpg", body, 0666)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true})
	w.Write(res)
}

func EditProductRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = io.ReadAll(r.Body)
	var id = r.URL.Query().Get("id")
	helpers.CheckNilErr(err)
	var product models.Product
	helpers.DecodeJson(body, &product)
	var res = productHelpers.EditProduct(product, id)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var res = productHelpers.DeleteProduct(id)
	os.Remove("public/images/" + id + ".jpg")
	w.Write(res)
}
