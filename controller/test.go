package controller

import (
	"net/http"
)

func TestRouter(w http.ResponseWriter, r *http.Request) {
	var name = r.URL.Query().Get("name")
	w.Write([]byte("Response from go api to " + name))
}
