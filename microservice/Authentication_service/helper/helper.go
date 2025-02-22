package helper

import (
	"authentication/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collections *mongo.Collection

func init() {
	// clientoptions := options.Client().ApplyURI("mongodb://admin:password@mongodb:27017")
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientoptions) // connection eshtablished
	if err != nil {
		log.Fatal(err)
	}
	collections = client.Database("Hotel").Collection("Users")
}

func SearchData(t model.Login_system) (model.User, bool) {
	curr, err := collections.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for curr.Next(context.Background()) {
		var a model.User
		err := curr.Decode(&a)
		if err != nil {
			log.Fatal(err)
		}
		if t.Username == a.Name && t.Password == a.Password {
			return a, true
		}
	}
	var b model.User
	return b, false
}
