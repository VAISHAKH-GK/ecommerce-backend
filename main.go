package main

import (
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/router"
	"github.com/rs/cors"
)

func main() {
	const port = ":9000"

	// starting http server
	http.ListenAndServe(port, cors.AllowAll().Handler(router.Router()))

	// disconnecting from mongodb
	defer databaseConnection.Client.Disconnect(databaseConnection.Ctx)
}
