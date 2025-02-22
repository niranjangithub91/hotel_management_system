package helper

import (
	"context"
	"fmt"
	"log"
	"user_management/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// clienopt := options.Client().ApplyURI("mongodb://admin:password@mongodb:27017")
	clienopt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clienopt)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("Hotel").Collection("Users")
}

func Add_Users(t model.User) {
	insert, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insert)
}
func Delete_users(t model.Person) {
	filter := bson.M{"name": t.Name}
	delete, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(delete)
}
