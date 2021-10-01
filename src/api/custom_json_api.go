package api

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/gorilla/mux"
)

func HandleCustomJson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var customjson datamodels.CustomJson
	err := json.NewDecoder(r.Body).Decode(&customjson)

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)

	}
	controller.CustomJsonSaveDetails(customjson)

	fmt.Print(r.Body)
	fmt.Println("CustomJson create Endpoint Hit\n")

}

func HandleGetCustomjsonDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := controller.GetCustomJsonDetails(params["id"])
	json.NewEncoder(w).Encode(result)
}

func HandleGetCustomjsonByProjectId(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	customJson := controller.GetCustomJsontByProject(params["projectId"])
	json.NewEncoder(w).Encode(customJson)

}