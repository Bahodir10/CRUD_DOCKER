package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"CRUD_DOCKER/models"
)

// RegisterUserRoutes registers user routes to the router
func RegisterUserRoutes(r *mux.Router, collection *mongo.Collection) {
	r.HandleFunc("/users", createUser(collection)).Methods("POST")
	r.HandleFunc("/users", getAllUsers(collection)).Methods("GET")
	r.HandleFunc("/users/{id}", getUser(collection)).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser(collection)).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser(collection)).Methods("DELETE")
}

func createUser(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		user.ID = primitive.NewObjectID()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func getAllUsers(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		var users []models.User
		for cursor.Next(ctx) {
			var user models.User
			cursor.Decode(&user)
			users = append(users, user)
		}

		json.NewEncoder(w).Encode(users)
	}
}

func getUser(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var user models.User
		err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func updateUser(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": user})
		if err != nil {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func deleteUser(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
		if err != nil {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
