package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Chatroom struct {
		Id          primitive.ObjectID `json:"id" bson:"_id, omitempty"`
		Name        string             `json:"name" bson:"name"`
		Owner       string             `json:"owner" bson:"owner"`
		Description string             `json:"description" bson:"description"`
		Members     []Member           `json:"members" bson:"members"`
		CreatedAt   time.Time          `json:"created_at" bson:"createdAt"`
		UpdatedAt   time.Time          `json:"updated_at" bson:"updatedAt"`
	}

	Member struct {
		Id      string    `json:"id" bson:"_id"`
		SinceAt time.Time `json:"since_at" bson:"sinceAt"`
	}
)
