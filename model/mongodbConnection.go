package model

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
	const uri = "mongodb://localhost:27017"
	Ctx = context.Background()
	var options = options.Client().ApplyURI(uri)
	Client, err = mongo.Connect(Ctx, options)
	checkNilErr(err)
	Db = Client.Database("ecommerce")
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
