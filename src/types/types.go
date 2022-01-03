package types

import "go.mongodb.org/mongo-driver/bson/primitive"


/*
* This class contains custom data type defeniotns and response wrappers
*
*/

/*
Type def to create JSON for HTML data generation
*/
type ReportScript struct {
	ReportName  string   `json:"reportName"`
	URL         string   `json:"url"`
	AccessToken string   `json:"accessToken"`
	Widgets     []Widget `json:"widgets"`
	FilterOptions  FilterOptions `json:"filterOptions"`
}

type FilterOptions  struct {
	FromDate    bool `json:"fromDate"`
	ToDate      bool `json:"toDate"`
	TenantIds   bool  `json:"tenantIds"`
	StageIds	bool  `json:"stageIds"`
	ItemIds		bool  `json:"itemIds"`

}
type Widget struct {

	WidgetType  string   `json:"widgetType"`
	WidgetName  string   `json:"widgetName"`
	Request     Request  `json:"request"`
	ColorScheme []string `json:"color"`
	XAxis       string   `json:"xAxis"`
	YAxis       string   `json:"yAxis"`
    Value      string  `json:"value"`
	Url string `json:"url"`
	Width string `json:"width"`
	Height string `json:"height"`

}

/*
 Type def , request struct for send request from report templates
 Tree  : JSON for query language
 Query : LDAL query
 Type  : Type of the JOSN , TDP/OTP/Normal
*/
type Request struct {
	// Tree  string `json:"tree"`
	Query string `json:"query"`
	Type  string `json:"type"`
}


type ReportTemplate struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ReportName     string             `json:"reportName"`
	Url            string             `json:"url"`
	ProjectId      string             `json:"projectId"`
	ReportTemplate string             `json:"template"`
}

/*
	Response wrapper for report query request
*/
type LDALReportQueryRequest struct {
	Tree  string `json:"tree"`
	Query string `json:"query"`
	Type  string `json:"type"`
}
