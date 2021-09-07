package datamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type LDALscript struct {
	SciptId     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ScriptName  string             `json:"scriptName"`
	ProjectID   string             `json:"projectId"`
	BoundStatus bool               `json:"boundStatus"`
	BoundedId   string             `json:"boundedId"`
	Content     string             `json:"content"`
}

type LDALscriptResult struct {
	SciptId string
	Result  string
}
