package helper

import (
	"context"
	"fmt"
	"log"
	"room_management/model"
	"room_management/userpb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	clientopt := options.Client().ApplyURI("mongodb://admin:password@mongodb:27017")
	client, err := mongo.Connect(context.Background(), clientopt)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("Hotel").Collection("Rooms")
}

func AddRooms_helper(t model.Room) {
	insert, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insert)
}

func DeleteRoom_helper(t userpb.SendDeleteRoomDetail) {
	filter := bson.M{"room_number": t.Roomnumber}
	delete, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(delete)
}
func UpdateFeatures_helper(t model.Room_feature_update) {
	filter := bson.M{"room_number": t.Room_numer}
	update := bson.M{"$set": bson.M{"features": t.Features}}
	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updated)
}
func UpdatePrice_helper(t model.Room_price_update) {
	fmt.Println("Hi")
	filter := bson.M{"room_number": t.Room_number}
	update := bson.M{"$set": bson.M{"price": t.Price}}
	updation, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updation)
}
