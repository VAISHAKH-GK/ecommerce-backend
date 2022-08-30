package controller

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/productHelpers"
)

func GetProductsRoute(w http.ResponseWriter, r *http.Request) {
	var numberOfProducts, err = strconv.Atoi(r.URL.Query().Get("number"))
	helpers.CheckNilErr(err)
	var res = productHelpers.GetAllProducts(numberOfProducts)
	w.Write(res)
}

func GetProductImageRotue(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var image, err = ioutil.ReadFile("public/images/" + id + ".jpg")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}
