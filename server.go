package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
