package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/config"
	"todo-api/models"
	"todo-api/views"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	var item models.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Convert user_id string to gocql.UUID
	userID, err := gocql.ParseUUID(item.UserID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}
	item.UserID = userID.String() // Convert back to string if needed

	if err := item.Save(); err != nil {
		views.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	views.SuccessResponse(w, item)
}

func GetTodoItems(w http.ResponseWriter, r *http.Request) {

	var requestBody struct {
		UserID string `json:"user_id"`
	}

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Check if user_id is provided
	if requestBody.UserID == "" {
		views.ErrorResponse(w, http.StatusBadRequest, "user_id is required")
		return
	}

	// Parse user_id to gocql.UUID
	userID, err := gocql.ParseUUID(requestBody.UserID)
	if err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	// Retrieve status and limit from URL query parameters
	status := r.URL.Query().Get("status")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}

	// Construct the CQL query
	query := `SELECT id, user_id, title, description, status, created, updated FROM todo WHERE user_id = ?`
	args := []interface{}{userID}

	if status != "" {
		query += ` AND status = ?`
		args = append(args, status)
	}

	query += ` ALLOW FILTERING`

	// Perform the query
	iter := config.Session.Query(query, args...).Iter()

	var items []models.TodoItem
	var item models.TodoItem
	for iter.Scan(&item.ID, &item.UserID, &item.Title, &item.Description, &item.Status, &item.Created, &item.Updated) {
		items = append(items, item)
	}

	if err := iter.Close(); err != nil {
		views.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Return the response with the todo items
	views.SuccessResponse(w, items)
}

func UpdateTodoItem(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	var item models.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	item.ID, _ = gocql.ParseUUID(idStr)

	if err := item.Update(); err != nil {
		views.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	views.SuccessResponse(w, item)
}

func DeleteTodoItem(w http.ResponseWriter, r *http.Request) {
	id, err := gocql.ParseUUID(mux.Vars(r)["id"])
	if err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, "Invalid UUID")
		return
	}

	if err := models.DeleteTodoItem(id); err != nil {
		views.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetTodoItemsByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["userID"]
	status := vars["status"]

	if userIDStr == "" {
		views.ErrorResponse(w, http.StatusBadRequest, "userID is required")
		return
	}

	if status == "" {
		views.ErrorResponse(w, http.StatusBadRequest, "status is required")
		return
	}

	userID, err := gocql.ParseUUID(userIDStr)
	if err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, "invalid UUID")
		return
	}

	items, err := models.GetTodoItemsByStatus(userID.String(), status)
	if err != nil {
		views.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	views.SuccessResponse(w, items)
}
