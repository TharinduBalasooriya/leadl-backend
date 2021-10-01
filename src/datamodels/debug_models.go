package datamodels

type DebugRequest struct {
	ProjectId  string `json:"project_id"`
	LDELScript string `json:"ldel_script"`
	LogFile    string `json:"log_file"`
}

type DebugResponse struct {
	Result string `json:"response"`
}

type LDALRequest struct {
	FileId string `json:"fileId"`
	Script string `json:"script"`
}
type LDALDebugRequest struct{
	Tree string  `json:"tree"`
	Query string `json:"query"`
	Type string  `json:"type"`
}