package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todo-api/models"
	"todo-api/views"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type TodoRequest struct {
	UserID string `json:"user_id"`
	Sort   string `json:"sort"`
}

func CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	var item models.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := gocql.ParseUUID(item.UserID)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}
	item.UserID = userID.String()

	if err := item.Save(); err != nil {
		views.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	views.SuccessResponse(w, item)
}

func GetTodoItems(w http.ResponseWriter, r *http.Request) {
	var req TodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.UserID == "" {
		views.ErrorResponse(w, http.StatusBadRequest, "user_id is required")
		return
	}

	status := r.URL.Query().Get("status")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit == 0 {
		limit = 10
	}

	items, err := models.GetTodoItems(req.UserID, status, limit)
	if err != nil {
		views.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

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

func GetSortTodoItems(w http.ResponseWriter, r *http.Request) {
	var req TodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		views.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.UserID == "" {
		views.ErrorResponse(w, http.StatusBadRequest, "user_id is required")
		return
	}

	limit := 10
	sort := "ASC"

	if req.Sort != "" {
		sort = strings.ToUpper(req.Sort)
		if sort != "ASC" && sort != "DESC" {
			views.ErrorResponse(w, http.StatusBadRequest, "Invalid sort order. Allowed values: ASC, DESC")
			return
		}
	}

	items, err := models.GetSortTodoItems(req.UserID, sort, limit)
	if err != nil {
		views.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if items == nil {
		views.SuccessResponse(w, "No todo items found")
	} else {
		views.SuccessResponse(w, items)
	}
}
