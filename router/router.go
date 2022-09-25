package router

import (
	"github.com/VAISHAKH-GK/ecommerce-backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	var router = mux.NewRouter()

	// user routes
	router.HandleFunc("/api/user/signup", controller.UserSignUpRoute).Methods("POST")
	router.HandleFunc("/api/user/login", controller.UserLoginRoute).Methods("POST")
	router.HandleFunc("/api/user/checklogin", controller.CheckLoginRoute).Methods("GET")
	router.HandleFunc("/api/user/getuser", controller.GetUserDataRoute).Methods("GET")
	router.HandleFunc("/api/user/logout", controller.UserLogoutRoute).Methods("GET")
	router.HandleFunc("/api/user/logout", controller.UserLogoutRoute).Methods("GET")
	router.HandleFunc("/api/user/addtocart", controller.AddToCartRoute).Methods("PUT")
	router.HandleFunc("/api/user/getcartproducts", controller.GetCartProductsRoute).Methods("GET")
	router.HandleFunc("/api/user/removefromcart", controller.RemoveFromCartRoute).Methods("DELETE")
	router.HandleFunc("/api/user/gettotal", controller.GetTotalPriceRoute).Methods("GET")
	router.HandleFunc("/api/user/placeorder", controller.PlaceOrderRoute).Methods("POST")
	router.HandleFunc("/api/user/getorders", controller.GetOrdersRotue).Methods("GET")
	router.HandleFunc("/api/user/getorderproducts", controller.GetOrderProductsRoute).Methods("GET")
	router.HandleFunc("/api/user/verifypayment", controller.VerifyPaymentRoute).Methods("POST")

	//admin routes
	router.HandleFunc("/api/admin/login", controller.AdminLoginRoute).Methods("POST")
	router.HandleFunc("/api/admin/addadmin", controller.AddAdminRoute).Methods("POST")
	router.HandleFunc("/api/admin/checklogin", controller.AdminCheckLoginRoute).Methods("GET")
	router.HandleFunc("/api/admin/getuser", controller.AdminGetUserDataRotue).Methods("GET")
	router.HandleFunc("/api/admin/logout", controller.AdminLogoutRoute).Methods("GET")
	router.HandleFunc("/api/admin/addproduct", controller.AddProductRoute).Methods("POST")
	router.HandleFunc("/api/admin/addproductimage", controller.AddProductImageRoute).Methods("POST")
	router.HandleFunc("/api/admin/updateproduct", controller.EditProductRoute).Methods("PUT")
	router.HandleFunc("/api/admin/deleteproduct", controller.DeleteProduct).Methods("DELETE")

	//public routes
	router.HandleFunc("/api/public/getproducts", controller.GetProductsRoute).Methods("GET")
	router.HandleFunc("/api/public/getproduct", controller.GetProductRoute).Methods("GET")
	router.HandleFunc("/api/public/getproductimage", controller.GetProductImageRotue).Methods("GET")
	router.HandleFunc("/api/public/searchproduct", controller.SearchProductRoute).Methods("GET")

	return router
}
