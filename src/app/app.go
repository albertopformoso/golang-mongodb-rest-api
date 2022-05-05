package app

import (
	"golang-mongo-rest-api/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	port := ":" + os.Getenv("PORT")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.GetPerson).Methods("GET")
	router.HandleFunc("/person", controllers.CreatePerson).Methods("POST")

	log.Fatal(http.ListenAndServe(port, router))
}
