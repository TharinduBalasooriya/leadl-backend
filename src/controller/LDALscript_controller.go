package controller

import (
	"fmt"
	"log"
	"os"

	fcllib "github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"

	"encoding/base64"
	"encoding/json"

	"github.com/google/uuid"
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

var result string

func ExecuteLDAL(scriptId string) string {

	requestId := uuid.New().String()
	var ldalDetails datamodels.LDALscript
	ldalDetails = ldalRepo.GetLDALScripts(scriptId)

	logFileDetails := logrepo.GetLogFileDetails(ldalDetails.BoundedId)
	service.Log_Download_LogFile(ldalDetails.BoundedId, requestId)
	service.Log_download_Script(ldalDetails.BoundedId, requestId)
	Config_LDEL_DEF(logFileDetails.LogFileName, requestId)
	service.Log_Execute_LDEL(requestId)
	decodedContent, err := base64.StdEncoding.DecodeString(ldalDetails.Content)
	if err != nil {
		log.Println("decode error:", err)

	}
	service.WriteToFile("localstorage/"+requestId, "LDAL_Script.txt", string(decodedContent))

	result := fcllib.NewFCLWrapper().GetLDALResult("localstorage/" + requestId + "/" + "Defs.txt")

	os.RemoveAll("localstorage/" + requestId)

	return result

}

func DebugLDAL(request datamodels.LDALDebugRequest) interface{} {
	requestID := uuid.New().String()
	decodedTree, err := base64.StdEncoding.DecodeString(request.Tree)
	if err != nil {
		log.Println("decode error:", err)

	}

	decodedQuery, err := base64.StdEncoding.DecodeString(request.Query)
	if err != nil {
		log.Println("decode error:", err)

	}

	service.WriteToFile("localstorage/"+requestID, "result.txt", string(decodedTree))
	service.WriteToFile("localstorage/"+requestID, "LDAL_Script.txt", string(decodedQuery))

	Config_LDEL_DEF("____", requestID)
	_ = fcllib.NewFCLWrapper().GetLDALResult("localstorage/" + requestID + "/" + "Defs.txt")
	dat, err := os.ReadFile("localstorage/" + requestID + "/Debug_Result.json")

	// an arbitrary json string
	jsonString := string(dat)

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)


	return jsonMap

}
