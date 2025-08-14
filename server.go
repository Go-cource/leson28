package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   primitive.ObjectID `json: "id" bson:"_id, omitempty"`
	Name string             `json: "name" bson: "name"`
	Age  int                `json: "age" bson: "age"`
}

var client *mongo.Client

func usersHandler(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("lesson28").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if r.Method == http.MethodGet {
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			fmt.Println(err)
		}
		defer cursor.Close(ctx)
		var users []User
		if err := cursor.All(ctx, &users); err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
	if r.Method == http.MethodPost {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			fmt.Println(err)
			return
		}
		result, err := collection.InsertOne(ctx, user)
		if err != nil {
			fmt.Println(err)
		}
		user.ID = result.InsertedID.(primitive.ObjectID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(ctx)

	http.HandleFunc("/users", usersHandler)
	fmt.Println("Server is started...")
	http.ListenAndServe(":8080", nil)
}
