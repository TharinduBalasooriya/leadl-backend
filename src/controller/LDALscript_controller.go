package controller

import (
	"log"
	"os"

	fcllib "github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/types"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"

	"encoding/base64"
	"encoding/json"

	"errors"

	"github.com/google/uuid"
)

var LDALscriptrepo repository.ScriptRepository
var cusjsonrepo repository.CustomJsontRepository

func ScriptSaveDetails(script datamodels.LDALscript) (interface{}, error) {
	results, err := service.Script_Save_Details(script)
	return results, err

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

//var result string

func ExecuteLDAL(scriptId string) (string, error) {

	var ldalDetails datamodels.LDALscript
	var result string
	requestId := uuid.New().String()

	ldalDetails = ldalRepo.GetLDALScripts(scriptId)

	if ldalDetails.LogQuery {
		logFileDetails := logrepo.GetLogFileDetails(ldalDetails.BoundedId)
		if len(logFileDetails.FileId) > 0 {
			service.Log_Download_LogFile(ldalDetails.BoundedId, requestId)
			//Download LDAL Script
			service.Log_download_Script(ldalDetails.BoundedId, requestId)
			Config_LDEL_DEF(logFileDetails.LogFileName, requestId)
			service.Log_Execute_LDEL(requestId)
			decodedContent, err := base64.StdEncoding.DecodeString(ldalDetails.Content)
			if err != nil {
				log.Println("decode error:", err)
				return "", err

			}
			service.WriteToFile("localstorage/"+requestId, "LDAL_Script.txt", string(decodedContent))

		} else {

			return "", errors.New("query Failed , log bind error")

		}
		result = fcllib.NewFCLWrapper().GetLDALResult("localstorage/" + requestId + "/" + "Defs.txt")

	} else {
		log.Println("Custom query execution started ...")
		if !ldalDetails.BoundStatus {
			return "Query is not bounded", nil
		} else {
			decodedContent, err := base64.StdEncoding.DecodeString(ldalDetails.Content)
			if err != nil {
				log.Println("decode error:", err)
				return "", err

			}
			service.WriteToFile("localstorage/"+requestId, "LDAL_Script.txt", string(decodedContent))

			service.DownloadCustomJSON(ldalDetails.BoundedId, requestId)
			//Check custom data types
			customJSONRequest := cusjsonrepo.GetCustomJson(ldalDetails.BoundedId)
			Config_LDEL_DEF("", requestId)
			if customJSONRequest.JsonType == "TDP" {
				result = fcllib.NewFCLWrapper().GetTDPResult("localstorage/" + requestId + "/" + "Defs.txt")
				//result="TDP"
			}else if customJSONRequest.JsonType == "OTP"{
				result =  fcllib.NewFCLWrapper().GetOTPResult("localstorage/" + requestId + "/" + "Defs.txt");

			} else if customJSONRequest.JsonType == "Normal" {
				//result = "Normal"
				result = fcllib.NewFCLWrapper().GetLogLDALResult("localstorage/" + requestId + "/" + "Defs.txt")
			} else {
				result = "Invalid custom json  Type"
			}

		}

	}
	os.RemoveAll("localstorage/" + requestId)
	log.Println(ldalDetails.BoundStatus)
	return result, nil
}

/**
* TODO : Implement OTP Processing capabilities
  to following methods
*
**/

func DebugLDAL(request datamodels.LDALDebugRequest) interface{} {

	requestID := uuid.New().String()
	var jsonMap map[string]interface{}
	var jsonString string
	if len(request.Query) > 0 && len(request.Tree) > 0 && len(request.Type) > 0 {

		service.WriteToFile("localstorage/"+requestID, "result.txt", string(request.Tree))
		service.WriteToFile("localstorage/"+requestID, "LDAL_Script.txt", string(request.Query))

		Config_LDEL_DEF("____", requestID)

		if request.Type == "Normal" {
			_ = fcllib.NewFCLWrapper().GetLogLDALResult("localstorage/" + requestID + "/" + "Defs.txt")
			data, err := os.ReadFile("localstorage/" + requestID + "/Debug_Result.json")

			if err != nil {
				jsonString = `{
				"variables": [
				{
					"dataType": "ERROR",
					"details": "Debug Failed",
					"name": "Debug Error"
					}
				]
			}`

			} else {
				jsonString = string(data)
			}

		} else if request.Type == "TDP" {

			_ = fcllib.NewFCLWrapper().GetTDPResult("localstorage/" + requestID + "/" + "Defs.txt")
			data, err := os.ReadFile("localstorage/" + requestID + "/Debug_Result.json")

			if err != nil {
				jsonString = `{
				"variables": [
				{
					"dataType": "ERROR",
					"details": "Debug Failed",
					"name": "Debug Error"
					}
				]
			}`

			} else {
				jsonString = string(data)
			}

		} else if request.Type == "OTP"{

			_ = fcllib.NewFCLWrapper().GetOTPResult("localstorage/" + requestID + "/" + "Defs.txt")
			data, err := os.ReadFile("localstorage/" + requestID + "/Debug_Result.json")

			if err != nil {
				jsonString = `{
				"variables": [
				{
					"dataType": "ERROR",
					"details": "Debug Failed",
					"name": "Debug Error"
					}
				]
			}`

			} else {
				jsonString = string(data)
			}

		}else {

			jsonString = `{
			"variables": [
			{
				"dataType": "ERROR",
				"details": "Invalid json type",
				"name": "SERVER ERROR"
				}
			]
		}`

		}

	} else {
		jsonString = `{
			"variables": [
			{
				"dataType": "ERROR",
				"details": "Request failed",
				"name": "REQUEST ERROR"
				}
			]
		}`
	}

	json.Unmarshal([]byte(jsonString), &jsonMap)
	os.RemoveAll("localstorage/" + requestID)
	return jsonMap

}

func ReportLDALRequest(request types.LDALReportQueryRequest) interface{} {

	requestID := uuid.New().String()
	var jsonMap map[string]interface{}
	var jsonString string
	if len(request.Query) > 0 && len(request.Tree) > 0 && len(request.Type) > 0 {

		service.WriteToFile("localstorage/"+requestID, "result.txt", string(request.Tree))
		service.WriteToFile("localstorage/"+requestID, "LDAL_Script.txt", string(request.Query))

		Config_LDEL_DEF("____", requestID)

		if request.Type == "Normal" || request.Type == "normal" {
			data := fcllib.NewFCLWrapper().GetLogLDALResult("localstorage/" + requestID + "/" + "Defs.txt")
			jsonString = data

		} else if request.Type == "TDP" || request.Type == "otp" {

			data := fcllib.NewFCLWrapper().GetTDPResult("localstorage/" + requestID + "/" + "Defs.txt")
			jsonString = data

		} else {

			jsonString = `{
			"variables": [
			{
				"dataType": "ERROR",
				"details": "Invalid json type",
				"name": "SERVER ERROR"
				}
			]
		}`

		}

	} else {
		jsonString = `{
			"variables": [
			{
				"dataType": "ERROR",
				"details": "Request failed",
				"name": "REQUEST ERROR"
				}
			]
		}`
	}

	json.Unmarshal([]byte(jsonString), &jsonMap)
	os.RemoveAll("localstorage/" + requestID)
	return jsonMap

}
