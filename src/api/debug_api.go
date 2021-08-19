package api

import (
	"encoding/json"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/gorilla/mux"
)

func HandelDebugLDEL(w http.ResponseWriter , r *http.Request ){

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := controller.GetLDELDebugResult(params["projectId"])
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
	 
	}


}

func HandleDebugProject(w http.ResponseWriter, r *http.Request){

	decoder := json.NewDecoder(r.Body)
	var debugRequest datamodels.DebugRequest
	err := decoder.Decode(&debugRequest)
	if err != nil{
		println(err.Error())
	}
	controller.SaveDebugProject(debugRequest)



}

