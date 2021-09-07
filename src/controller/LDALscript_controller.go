package controller

import (
	"fmt"
	"log"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var LDALscriptrepo repository.ScriptRepository

func ScriptSaveDetails(script datamodels.LDALscript) {

	results, err := service.Script_Save_Details(script)

	if err != nil {
		log.Fatal(err)

	}

	id := results.(primitive.ObjectID)
	fmt.Println("Successfully inserted" + id.String())

}

func UpdateScript(script datamodels.LDALscript) interface{} {

	results := LDALscriptrepo.UpadteLDALScript(script)

	return results
}

func GetScriptDetails(scriptId string) datamodels.LDALscript {
	result := LDALscriptrepo.GetLDALScripts(scriptId)
	return result
}

func GetScriptByProject(projectId string) []datamodels.LDALscript {

	scriptList := LDALscriptrepo.GetScriptsByProjectID(projectId)

	return scriptList
}
