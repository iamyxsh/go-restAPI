package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type persons struct {
	ID primitive.ObjectID `json: "_id, omitempty" bson: "_id, omitempty`
	FirstName string `json: "firstName, omitempty" bson: "firstName, omitempty`
	LastName string `json: "lastName, omitempty" bson: "lastName, omitempty`
	Hobbies []string `json: "hobbies" bson: "hobbies`
} 

var client mongo.Client

func createPerson(res http.ResponseWriter, req *http.Request) {
	 res.Header().Add("content-type", "application/json")
	 var person persons

	 json.NewDecoder(req.Body).Decode(&person)
	 collection := client.Database("GoLang").Collection("people")
	 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	 defer cancel()
	 result, err := collection.InsertOne(ctx, person)

	 if err != nil {
		 log.Fatal(err)
	 } 

	 json.NewEncoder(res).Encode(result)

}

func main ()  {
	fmt.Println("Hello")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://yxsh:Qwerty12345@portfolio-apps.zrui0.mongodb.net/GoLang?retryWrites=true&w=majority") )

	if err != nil{
		log.Fatal(err)
	} 

	if client != nil {
		fmt.Println("Connected to database.")
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/person", createPerson).Methods("POST")

	http.ListenAndServe(":5000", router)

}