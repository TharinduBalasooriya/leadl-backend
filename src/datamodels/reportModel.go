package datamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ReportName string `json:"reportName"`
	ProjectId string `json:"projectId"`
	Widgets []Widget `json:"widgets"`
}
type Widget struct {
	ID 		 string `json:"id"`
	Type     string `json:"type"`
	ScriptId string `json:"scriptId"`
	X        string `json:"x"`
	Y        string `json:"y"`
	Value    string `json:"value,omitempty"`
}
