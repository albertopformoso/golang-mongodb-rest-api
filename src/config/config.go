package config

import (
	"encoding/json"
	"net/http"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (collection *mongo.Collection) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	// Connect to Mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection = client.Database("go-mongo").Collection("People")

	return
}

// ErrorResponse : This is error model
type ErrorResponse struct {
	StatusCode int `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare the error model
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse {
		ErrorMessage: err.Error(),
		StatusCode: http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
