package helper

import (
	"context"
	"customer/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collections *mongo.Collection

func init() {
	// clientopt := options.Client().ApplyURI("mongodb://admin:password@mongodb:27017")
	clientopt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientopt)
	if err != nil {
		log.Fatal(err)
	}
	collections = client.Database("Hotel").Collection("Customers")
}

func Add_Customers_helper(t model.Customer) {
	insert, err := collections.InsertOne(context.Background(), t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insert)
}
