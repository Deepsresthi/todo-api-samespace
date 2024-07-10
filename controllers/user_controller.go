package controllers

import (
	"encoding/json"
	"net/http"
	"todo-api/models"

	"github.com/gocql/gocql"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the user ID
	userID := gocql.TimeUUID()

	// Convert UUID to string before assigning to user.UserID
	user.UserID = userID

	if err := user.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created user with HTTP status 201 Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}