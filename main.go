package main

import (
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/model"
	"github.com/VAISHAKH-GK/ecommerce-backend/router"
	"github.com/rs/cors"
)

func main() {
	const port = ":9000"

  // starting http server
	http.ListenAndServe(port, cors.AllowAll().Handler(router.Router()))
  
  // disconnecting from mongodb
	defer model.Client.Disconnect(model.Ctx)
}
