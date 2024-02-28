package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents a user in the database.
type User struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
}

// ConnectToDB creates a MongoDB client and connects to the database.
func ConnectToDB() *mongo.Client {
	uri := "mongodb+srv://jguida:Vidatirana12*@cluster0.k0lgz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0" // Make sure to replace this with your actual URI
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// AddUser adds a new user to the database.
func AddUser(client *mongo.Client, w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("disponibilidade").Collection("jgos") // Replace "your_database_name" with your actual database name
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	client := ConnectToDB() // Connect to MongoDB Atlas

	r := mux.NewRouter()
	r.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		AddUser(client, w, r)
	}).Methods("POST")

	fmt.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
