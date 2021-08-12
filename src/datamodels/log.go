package datamodels

type Log struct {
	Username    string `json:"username"`
	ProjectId   string `json:"projectId"`
	LogFileName string `json:"logfilename"`
	LastUpdate  string `json:"lastupdate"`
	FileId      string `json:"fileId"`
}

type Log_Update struct {
	FileId  string `json:"fileId"`
	Content string `json:"content"`
}
