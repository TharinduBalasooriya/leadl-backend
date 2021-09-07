package api

import (
	"encoding/json"
	"fmt"
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
	controller.ScriptSaveDetails(script)

	fmt.Print(r.Body)
	fmt.Println("script create Endpoint Hit\n")

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