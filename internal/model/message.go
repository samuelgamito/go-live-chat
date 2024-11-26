package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	From    string             `json:"from" bson:"from"`
	To      string             `json:"to" bson:"to"`
	Content string             `json:"content" bson:"content"`
	Type    string             `json:"type" bson:"type"`
}
