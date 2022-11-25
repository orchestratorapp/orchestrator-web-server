package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Student struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Document    string `json:"document"`
	PhoneNumber string `json:"phone_number"`
}

// Handles the root ("/") path
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola!")
}

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
	w.Write(response)
	LogInfo(http.StatusCreated, r.URL.Path, time.Now())
}
