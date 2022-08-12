package router

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	var router = mux.NewRouter()

	// test route
	router.HandleFunc("/api/test", controller.TestRouter)

	// user routes
	router.HandleFunc("/api/user/signup", controller.UserSignUpRoute).Methods("POST")
	router.HandleFunc("/api/user/login", controller.UserLoginRoute).Methods("POST")
	router.HandleFunc("/api/user/checklogin", controller.CheckLoginRoute).Methods("GET", "POST", "PUT", "PATCH")
	router.HandleFunc("/api/user/getuser", controller.GetUserDataRoute).Methods("GET")

	//admin routes
	router.HandleFunc("/api/admin/login", controller.AdminLoginRoute).Methods("POST")
	router.HandleFunc("/api/admin/addadmin", controller.AddAdminRoute).Methods("POST")

	return router
}
