package controller

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/model"
	"net/http"
)

func TestRouter(w http.ResponseWriter, r *http.Request) {
	var name = r.URL.Query().Get("name")
	go func() {
		var data = map[string]interface{}{"name": name}
		model.Db.Collection("test").DeleteOne(model.Ctx, data)
	}()
	w.Write([]byte("Response from go api"))
}
