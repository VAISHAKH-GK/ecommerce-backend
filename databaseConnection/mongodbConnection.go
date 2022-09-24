package databaseConnection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx context.Context
var Client *mongo.Client
var Db *mongo.Database
var err error

func init() {
	// connecting uri for mongodb
	const uri = "mongodb://localhost:27017"
	// context for mongodb
	Ctx = context.Background()
	// creating options for mogndb connection and applying uri
	var options = options.Client().ApplyURI(uri)
	// connecting to mognodb
	Client, err = mongo.Connect(Ctx, options)
	// checking for any error while creating connection to mongodb
	checkNilErr(err)
	// storing db
	Db = Client.Database("ecommerce")
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Disconnect() {
	Client.Disconnect(Ctx)
}
