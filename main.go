package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"CRUD_DOCKER/database"
	"CRUD_DOCKER/handlers"
)

func main() {
	// Connect to MongoDB
	client, collection := database.ConnectMongoDB()
	defer client.Disconnect(database.Ctx)

	// Set up the router
	r := mux.NewRouter()

	// Register routes
	handlers.RegisterUserRoutes(r, collection)

	// Start the server
	log.Println("Server is running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", r))
}
