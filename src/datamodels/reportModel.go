package datamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
  Report table contains all the report configueration for a report
*/
type Report struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ReportName    string             `json:"reportName"`
	Url           string             `json:"url"`
	ProjectId     string             `json:"projectId"`
	Widgets       []Widget           `json:"widgets"`
	AccessToken   string             `json:"accessToken"`
	Headers       []Header           `json:"headers"`
	POSTParams    []POSTParam        `json:"postParams"`
	FilterOptions FilterOptions      `json:"filterOptions"`
}

/*
to store filter options value for report requests
*/
type FilterOptions struct {
	FromDate  bool `json:"fromDate"`
	ToDate    bool `json:"toDate"`
	TenantIds bool `json:"tenantIds"`
	StageIds  bool `json:"stageIds"`
	ItemIds   bool `json:"itemIds"`
}

/*
to store customized HTTP header value for report requests
*/
type Header struct {
	Name       string `json:"name"`
	JsFunction string `json:"jsFunction"`
}

/**
to store customized HTTP POST requets paramters
*/
type POSTParam struct {
	Name       string `json:"name"`
	JsFunction string `json:"jsFunction"`
}

type Widget struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	ScriptId  string   `json:"scriptId"`
	X         string   `json:"x"`
	Y         string   `json:"y"`
	Text      string   `json:"text"`
	Styles    Style    `json:"style"`
	BarCharts BarChart `json:"barGraph"`
	Url       string   `json:"url"`
	Width     string   `json:"width"`
	Height    string   `json:"height"`
}

type BarChart struct {
	Colors []string `json:"colors"`
	XAxis  string   `json:"xAxis"`
	YAxis  string   `json:"yAxis"`
}

//Style : styling details of widget
type Style struct {
	Fill                string `json:"fill"`
	TextBackgroundColor string `json:"textBackgroundColor"`
	BorderColor         string `json:"borderColor"`
	FontFamily          string `json:"fontFamily"`
	FontSize            int64  `json:"fontSize"`
	FontStyle           string `json:"fontStyle"`
	FontWeight          string `json:"fontWeight"`
	Stroke              string `json:"stroke"`
	TextAlign           string `json:"textAlign"`
	Value               string `json:"value,omitempty"`
}
