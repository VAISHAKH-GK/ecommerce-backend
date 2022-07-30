package router

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	var router = mux.NewRouter()

	router.HandleFunc("/api/test", controller.TestRouter)

	router.HandleFunc("/api/user/signup", controller.SignUp).Methods("POST")
	router.HandleFunc("/api/user/login", controller.Login).Methods("POST")

	return router
}
