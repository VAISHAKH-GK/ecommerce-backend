package main

import (
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/databaseConnection"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/router"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	const port = ":9000"

  var err = godotenv.Load(".env")
  helpers.CheckNilErr(err)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	})

	// starting http server
	http.ListenAndServe(port, c.Handler(router.Router()))

	// disconnecting from mongodb
	defer databaseConnection.Client.Disconnect(databaseConnection.Ctx)
}
