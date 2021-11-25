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
	Text     string `json:"text"`
	Styles   Style `json:"style"`
}


type Style struct {

  Fill  string `json:"fill"`;
  TextBackgroundColor  string `json:"textBackgroundColor"`;
  BorderColor  string `json:"borderColor"`;
  FontFamily  string `json:"fontFamily"`;
  FontSize int64 `json:"fontSize"`;
  FontStyle string `json:"fontStyle"`;
  FontWeight string `json:"fontWeight"`;
  Stroke string `json:"stroke"`;
  TextAlign string `json:"textAlign"`;

	Value    string `json:"value,omitempty"`
}

//HTML Template
type ReportTemplate struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ReportName string `json:"reportName"`
	ProjectId string `json:"projectId"`
	ReportTemplate string `json:"template"`


}