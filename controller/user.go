package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/VAISHAKH-GK/ecommerce-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var db = model.Db
var ctx = model.Ctx

func SignUp(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	checkNilErr(err)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	password, err := json.Marshal(data["password"])
	checkNilErr(err)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	checkNilErr(err)
	var hashedPassword = string(hash)
	data["password"] = hashedPassword
	response, err := db.Collection("user").InsertOne(ctx, data)
	checkNilErr(err)
	id, err := json.Marshal(response.InsertedID)
	checkNilErr(err)
	w.Write(id)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	checkNilErr(err)
	var data map[string]interface{}
	var user map[string]interface{}
	err = json.Unmarshal(body, &data)
	db.Collection("user").FindOne(ctx, bson.M{"email": data["email"]}).Decode(&user)
	var hashedPassword = user["password"].(string)
	password, err := json.Marshal(data["password"])
	checkNilErr(err)
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	var details map[string]interface{}
	if err == nil {
		details = map[string]interface{}{"status": true}
		res, err := json.Marshal(details)
		checkNilErr(err)
		w.Write(res)
	} else {
		details = map[string]interface{}{"status": false, "reason": "Wrong Password"}
		res, err := json.Marshal(details)
		checkNilErr(err)
		w.Write(res)
	}
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
