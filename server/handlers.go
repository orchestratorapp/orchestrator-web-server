package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type Student struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Document    string `json:"document"`
	PhoneNumber string `json:"phone_number"`
}

// Handles the root ("/") path. This function serves as an
// example handler, and should be deleted as you write your
// own code :)
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal("Hello, world!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError(http.StatusInternalServerError, r.URL.Path, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
	LogInfo(http.StatusOK, r.URL.Path, time.Now())
}

// Handles the post ("/post") path. This function serves as an
// example handler, and should be deleted as you write your
// own code :)
func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var student Student
	err := decoder.Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		LogError(http.StatusBadRequest, r.URL.Path, err)
		return
	}
	response, err := json.Marshal(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError(http.StatusInternalServerError, r.URL.Path, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	LogInfo(http.StatusCreated, r.URL.Path, time.Now())
}
