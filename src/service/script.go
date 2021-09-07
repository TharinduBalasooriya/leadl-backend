package service

import (

	//Importing file storage utility

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
)

var scriptrepo repository.ScriptRepository

//Save Log Details in mongo db
func Script_Save_Details(script datamodels.LDALscript) (interface{}, error) {
	resultID, err := scriptrepo.SaveLDALScript(script)
	return resultID, err

}
