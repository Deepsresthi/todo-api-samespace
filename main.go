package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/handlers"

	"todo-api/config"
	"todo-api/routes"
)

func main() {
	// Initialize ScyllaDB configuration
	scyllaConfig := config.ScyllaDBConfig{
		Hosts:       []string{"127.0.0.1"}, // Update with your ScyllaDB host(s)
		Keyspace:    "todo",
		Consistency: gocql.Quorum,
	}
	config.InitScyllaDB(scyllaConfig)
	defer config.CloseScyllaDB()

	r := routes.RegisterRoutes()

	
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	
	srv := &http.Server{
		Handler:      handlers.CORS(headers, methods, origins)(r),
		Addr:         ":8080", 
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started at", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
