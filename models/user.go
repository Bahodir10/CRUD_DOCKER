package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User defines the structure for the user model
type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}
