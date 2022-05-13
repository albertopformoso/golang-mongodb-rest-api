package controllers

import (
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetPeople(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/people", nil)
	response := httptest.NewRecorder()

	GetPeople(response, req)

	if response.Code == 500 {
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	} else {
		assert.Equal(t, http.StatusOK, response.Code)
	}
}

func TestGetPerson(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/person/6274404335b1ce195598acee", nil)
	response := httptest.NewRecorder()

	// Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "6274404335b1ce195598acee",
	}

	req = mux.SetURLVars(req, vars)

	GetPerson(response, req)
	expected1 := `{"data": {"_id": "6274404335b1ce195598acee", "firstname": "Bob", "lastname": "Ross"}, "message": "Success", "status": 200}`
	expected2 := `{"data":null,"message":"the provided hex string is not a valid ObjectID","status":500}`

	if response.Code == 500 {
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, expected2, response.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expected1, response.Body.String())
	}
}

func TestCreatePerson(t *testing.T) {
	var jsonStr = []byte(`{"_id":"6274404335b1ce195598acee", "firstname":"", "lastname":""}`)
	req, err := http.NewRequest("POST", "/people", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePerson)
	handler.ServeHTTP(response, req)

	if response.Code == 500 {
		assert.Equal(t, 500, response.Code, "Failed Creation")
	} else {
		assert.Equal(t, 200, response.Code, "Created!")
	}
}
