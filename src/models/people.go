package models

import (
	"golang-mongo-rest-api/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mongo database connection
var collection = config.ConnectDB()

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type People []Person
