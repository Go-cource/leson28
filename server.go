package main

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID   primitive.ObjectID `json: "id" bson:"_id, omitempty"`
	Name string             `json: "name" bson: "name"`
	Age  int                `json: "age" bson: "age"`
}

var client *mongo.Client

func main() {
	

	http.HandleFunc("/users", usersHandler)
	fmt.Println("Server is started...")
	http.ListenAndServe(":8080", nil)
}
