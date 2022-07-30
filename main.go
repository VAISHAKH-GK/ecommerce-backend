package main

import (
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/model"
	"github.com/VAISHAKH-GK/ecommerce-backend/router"
	"github.com/rs/cors"
)

func main() {
	const port = ":9000"

	http.ListenAndServe(port, cors.AllowAll().Handler(router.Router()))
	model.Client.Disconnect(model.Ctx)
}
