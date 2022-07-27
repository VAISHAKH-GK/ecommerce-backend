package controller

import (
	"net/http"
)

func IndexRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Response from go api"))
}
