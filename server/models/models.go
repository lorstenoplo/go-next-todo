package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoList struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Task      string             `json:"task" bson:"title"`
	Completed bool               `json:"completed" bson:"completed"`
}
