package routes

import (
	"todo-api/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	//to-do list
	r.HandleFunc("/todo", controllers.CreateTodoItem).Methods("POST")
	r.HandleFunc("/todo", controllers.GetTodoItems).Methods("GET")
	r.HandleFunc("/todo/{id}", controllers.UpdateTodoItem).Methods("PUT")
	r.HandleFunc("/todo/{id}", controllers.DeleteTodoItem).Methods("DELETE")
	return r
}
