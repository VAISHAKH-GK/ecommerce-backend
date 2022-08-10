package main

import (
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/router"
	"github.com/rs/cors"
)

func main() {
	const port = ":9000"

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	// starting http server
	http.ListenAndServe(port, c.Handler(router.Router()))

	// disconnecting from mongodb
	defer databaseConnection.Client.Disconnect(databaseConnection.Ctx)
}
