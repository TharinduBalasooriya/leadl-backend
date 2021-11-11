package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/gorilla/mux"
)

func HandleScripts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var script datamodels.LDALscript
	err := json.NewDecoder(r.Body).Decode(&script)

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)

	}
	result,err :=controller.ScriptSaveDetails(script)

	if err != nil{
		log.Println("Script create failed")
		json.NewEncoder(w).Encode("{\"Message\":\"Query Creation Failed\"}")
	}else{
		json.NewEncoder(w).Encode(result)
	}



}

func HandleGetScriptDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := controller.GetScriptDetails(params["id"])
	json.NewEncoder(w).Encode(result)
}

func HandleGetScriptsByProjectId(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	scripts := controller.GetScriptByProject(params["projectId"])
	fmt.Print(scripts)
	json.NewEncoder(w).Encode(scripts)

}


func HandleUpdateScripts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var script datamodels.LDALscript
	err := json.NewDecoder(r.Body).Decode(&script)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}
	result := controller.UpdateScript(script)
	fmt.Print(result)
	json.NewEncoder(w).Encode(result)

}

func HandelDebugLDAL(w http.ResponseWriter, r *http.Request){

	decoder := json.NewDecoder(r.Body)
	var debugRequest datamodels.LDALDebugRequest
	err := decoder.Decode(&debugRequest)
	if(err != nil){
		log.Println(err)
	}
	result := controller.DebugLDAL(debugRequest)

	json.NewEncoder(w).Encode(result)

	

}