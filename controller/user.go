package controller

import (
	"io"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/helpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/paymentHelpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/productHelpers"
	"github.com/VAISHAKH-GK/ecommerce-backend/helpers/userHelpers"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// post request on /api/user/signup
func UserSignUpRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res = userHelpers.DoUserSignUp(body)
	w.Write(res)
}

// post request on /api/user/login
func UserLoginRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var body, err = io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var res, userId = userHelpers.DoUserLogin(body)
	const twentyMinute = 60 * 1 * 20
	session, err := store.Get(r, "user")
	helpers.CheckNilErr(err)
	session.Options.MaxAge = twentyMinute
	session.Values["userId"] = userId.Hex()
	session.Values["isLoggedIn"] = true
	err = session.Save(r, w)
	helpers.CheckNilErr(err)
	w.Write(res)
}

// request on /api/user/checklogin
func CheckLoginRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	session, err := store.Get(r, "user")
	helpers.CheckNilErr(err)
	var res = userHelpers.CheckUserLogin(session)
	w.Write(res)
}

// get request on /api/user/getuser
func GetUserDataRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	session, err := store.Get(r, "user")
	helpers.CheckNilErr(err)
	var res = userHelpers.GetUserData(session)
	w.Write(res)
}

// get request on /api/user/logout
func UserLogoutRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "user")
	helpers.CheckNilErr(err)
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	helpers.CheckNilErr(err)
	var res = helpers.EncodeJson(map[string]interface{}{"status": true})
	w.Write(res)
}

// put request on route /api/user/addtocart
func AddToCartRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "user")
	helpers.CheckNilErr(err)
	if !userHelpers.CheckLogin(session) {
		w.Write(userHelpers.NotLoggedInResponse())
		return
	}
	var userId = userHelpers.GetUserId(session)
	body, err := io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var data map[string]interface{}
	helpers.DecodeJson(body, &data)
	var count = int(data["count"].(float64))
	productId, err := primitive.ObjectIDFromHex(data["productId"].(string))
	helpers.CheckNilErr(err)
	var res = productHelpers.AddProductToCart(userId, productId, count)
	w.Write(res)
}

// get request on /api/user/getcartproducts
func GetCartProductsRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "user")
	helpers.CheckNilErr(err)
	if session.Values["isLoggedIn"] != true {
		var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Not Logged In"})
		w.Write(res)
		return
	}
	userId, err := primitive.ObjectIDFromHex(session.Values["userId"].(string))
	helpers.CheckNilErr(err)
	var res = productHelpers.GetCartProducts(userId)
	w.Write(res)
}

// delete request on /api/user/removefromcart
func RemoveFromCartRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "user")
	helpers.CheckNilErr(err)
	if session.Values["isLoggedIn"] != true {
		var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Not Logged In"})
		w.Write(res)
		return
	}
	userId, err := primitive.ObjectIDFromHex(session.Values["userId"].(string))
	helpers.CheckNilErr(err)
	productId, err := primitive.ObjectIDFromHex(r.URL.Query().Get("productId"))
	helpers.CheckNilErr(err)
	var res = productHelpers.RemoveCartProduct(userId, productId)
	w.Write(res)
}

// getrequest on /api/user/gettotal
func GetTotalPriceRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "user")
	helpers.CheckNilErr(err)
	userHelpers.CheckLogin(session)
	userId, err := primitive.ObjectIDFromHex(session.Values["userId"].(string))
	helpers.CheckNilErr(err)
	var res = productHelpers.GetTotalAmount(userId)
	w.Write(res)
}

// post request on /api/usre/placeorder
func PlaceOrderRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "user")
	helpers.CheckNilErr(err)
	if !userHelpers.CheckLogin(session) {
		var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Not LoggedIn"})
		w.Write(res)
	}
	body, err := io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var data map[string]interface{}
	helpers.DecodeJson(body, &data)
	var userId = userHelpers.GetUserId(session)
	var res = userHelpers.PlaceOrder(data, userId)
	w.Write(res)
}

// get request on /api/user/getorders
func GetUserOrdersRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "user")
	helpers.CheckNilErr(err)
	if !userHelpers.CheckLogin(session) {
		var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Not LoggedIn"})
		w.Write(res)
		return
	}
	var userId = userHelpers.GetUserId(session)
	var order = userHelpers.GetOrders(userId)
	var res = helpers.EncodeJson(order)
	w.Write(res)
}

// get request on /api/user/getorderproducts
func GetOrderProductsRoute(w http.ResponseWriter, r *http.Request) {
	var store = sessions.NewCookieStore([]byte("ecommerce"))
	var session, err = store.Get(r, "user")
	helpers.CheckNilErr(err)
	if !userHelpers.CheckLogin(session) {
		var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Not LoggedIn"})
		w.Write(res)
		return
	}
	orderId, err := primitive.ObjectIDFromHex(r.URL.Query().Get("orderId"))
	helpers.CheckNilErr(err)
	var products = productHelpers.GetOrderProducts(orderId)
	w.Write(helpers.EncodeJson(products))
}

// post route on /api/user/verifypayment
func VerifyPaymentRoute(w http.ResponseWriter, r *http.Request) {
	var body, err = io.ReadAll(r.Body)
	helpers.CheckNilErr(err)
	var data map[string]interface{}
	helpers.DecodeJson(body, &data)
	orderId, err := primitive.ObjectIDFromHex(data["orderId"].(string))
	helpers.CheckNilErr(err)
	if paymentHelpers.VerifyOnlinePayment(data["payment"].(map[string]interface{})) {
		userHelpers.ChangeOrderStatus(orderId)
		var res = helpers.EncodeJson(map[string]interface{}{"status": true})
		w.Write(res)
		return
	}
	var res = helpers.EncodeJson(map[string]interface{}{"status": false, "reason": "Payment Not verified"})
	w.Write(res)
}
