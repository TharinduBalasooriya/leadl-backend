package datamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomJson struct {
	CustomJsonId     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomJsonName  string             `json:"customJsonName"`
	ProjectID   string             `json:"projectId"`
	JsonType   string             `json:"jsonType"`
	Content     string             `json:"content"`
}