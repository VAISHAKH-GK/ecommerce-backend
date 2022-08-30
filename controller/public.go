package controller

import (
	"io/ioutil"
	"net/http"
)

func DisplayImageRotue(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var image, err = ioutil.ReadFile("public/images/" + id + ".jpg")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}
