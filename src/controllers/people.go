package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang-mongo-rest-api/config"
	"golang-mongo-rest-api/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/matryer/respond.v1"
)

var collection = config.ConnectDB()

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Golang REST API with MongoDB")
}

func GetPeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var people models.People

	// Set Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Cancel to prevent memory leakage
	defer cancel()

	// Query data
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}

	defer cursor.Close(ctx)

	// Loop through person
	for cursor.Next(ctx) {
		var person models.Person
		cursor.Decode(&person)
		people = append(people, person)
	}

	// Handle error
	if err := cursor.Err(); err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}

	// Handle Success
	data := map[string]interface{}{"data": people, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)
}

func GetPerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	// Request parameters
	params := mux.Vars(request)
	// Convert id to usable mongodb object _id
	id, errorId := primitive.ObjectIDFromHex(params["id"])
	if errorId != nil {
		data := map[string]interface{}{"data": nil, "message": errorId.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}

	var person models.Person

	// Set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Cancel to prevent memory leakage
	defer cancel()

	// Query the model
	err := collection.FindOne(ctx, models.Person{ID: id}).Decode(&person)

	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}

	// Handle Success data
	data := map[string]interface{}{"data": person, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)
}

func CreatePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var person models.Person

	// Decode the body request parameters
	json.NewDecoder(request.Body).Decode(&person)

	// Validation for empty fields
	validate := validator.New()
	err := validate.Struct(person)

	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}

	// Set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//cancel to prevent memory leakage
	defer cancel()

	// insert our people model
	result, err := collection.InsertOne(ctx, person)

	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}

	data := map[string]interface{}{"data": result, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)
}
