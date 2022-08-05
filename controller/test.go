package controller

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"net/http"
)

func TestRouter(w http.ResponseWriter, r *http.Request) {
	var name = r.URL.Query().Get("name")
	go func() {
		var data = map[string]interface{}{"name": name}
		databaseConnection.Db.Collection("test").DeleteOne(databaseConnection.Ctx, data)
	}()
	w.Write([]byte("Response from go api"))
}
