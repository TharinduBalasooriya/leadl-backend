package api

//API to handle main log operation
//Get file name , Get content

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/gorilla/mux"
)

/*
	Read the content of log file
*/


func GetLogFileContentv2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//logs := controller.LogGetFileContent(params["user"], params["project"], params["logfileName"])
	log := controller.LogGetFileContentv2(params["fileId"])
	json.NewEncoder(w).Encode(log)
}



func GetLogListByProjectID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logs := controller.GetLogListByProjectID(params["id"])
	fmt.Println(params["id"])
	json.NewEncoder(w).Encode(logs)

}

func HandleActiavetLogFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result := controller.GetToActiveDir(params["fileId"])
	json.NewEncoder(w).Encode(result)

}
func HandleLogFileUpload(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("logFile")
	userName := r.FormValue("userName")
	projectName := r.FormValue("projectName")
	fileName := r.FormValue("fileName")
	fileId := r.FormValue("fileId")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//aws upload path
	fullFilePath := "logs/" + userName + "/" + projectName + "/" + fileName

	controller.LogSaveDetails(userName, projectName, fileName, fileId)
	controller.LogUploadFiles(fullFilePath, file)

	//controller.Config_LDEL_DEF(fileName, fileId)

}

type Update struct {
	UserName    string `json:"userName"`
	ProjectName string `json:"project"`
	Data        string `json:"data"`
}

func HandleFileUpdates(w http.ResponseWriter, r *http.Request) {
	var newupdate Update
	_ = json.NewDecoder(r.Body).Decode(&newupdate)
	controller.HandleUpdateData(controller.Update(newupdate))
	json.NewEncoder(w).Encode(newupdate)
}

func HandleSrciptUpload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("File Upload Endpoint Hit")
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	fileId := r.FormValue("fileId")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("localstorage/"+fileId+"/script.txt", fileBytes, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func HandleInvokeELInterpreter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, _ := controller.ExecuteLDEL(params["fileId"])
	json.NewEncoder(w).Encode(result)
}

func HandleInvokeELInterpreterGetJSON(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_, result := controller.ExecuteLDEL(params["fileId"])
	json.NewEncoder(w).Encode(result)
}

func HandleLogFileUpdate(w http.ResponseWriter, r *http.Request) {
	var logfileDetails datamodels.Log_Update
	err := json.NewDecoder(r.Body).Decode(&logfileDetails)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	result := controller.LogUpdateFile(logfileDetails)
	fmt.Fprintln(w, result)
}

func HandleExecuteLDAL(w http.ResponseWriter, r *http.Request) {
	//var ldalRequest datamodels.LDALRequest
	params := mux.Vars(r)

	result,err := controller.ExecuteLDAL(params["scriptId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		var LDALResult datamodels.LDALscriptResult
		LDALResult.SciptId = params["scriptId"]
		LDALResult.Result = result
		json.NewEncoder(w).Encode(LDALResult)
	}
	log.Println(result)
}
