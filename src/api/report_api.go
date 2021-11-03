package api

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/gorilla/mux"
)

func HandleCreateReport(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var report datamodels.Report
	err := json.NewDecoder(r.Body).Decode(&report)

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)

	}
	result,err := controller.ReportSaveDetails(report)

	if(err != nil){
		json.NewEncoder(w).Encode(err.Error())
	}else{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}

}

func HandleGetReportById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := controller.GetReportById(params["id"])
	json.NewEncoder(w).Encode(result)
}

func HandleGetReportsByProjectId(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	reports := controller.GetReportsByProject(params["projectId"])
	json.NewEncoder(w).Encode(reports)

}


func HandleUpdateReport(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var report datamodels.Report
	err := json.NewDecoder(r.Body).Decode(&report)

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)

	}
	result,err := controller.UpdateReport(report)

	if(err != nil){
		json.NewEncoder(w).Encode(err.Error())
	}else{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}

}

func HandleDeleteReportById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result,err := controller.DeleteReport(params["id"])
	if(err != nil){
		json.NewEncoder(w).Encode(err.Error())
	}else{
		json.NewEncoder(w).Encode(result)
	}
	
}


