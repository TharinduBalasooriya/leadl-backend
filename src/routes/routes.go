package routes

import (
	"fmt"

	"net/http"

	//"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/middleware"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/websocket"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/api"
	"github.com/gorilla/mux"
)

func LogRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	apiRoutes := router.PathPrefix("/api").Subrouter().StrictSlash(true)

	//Get All Log files
	router.HandleFunc("/ws", websocket.WSPage).Methods("GET")
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "HomeRoute")

	})

	//apiRoutes.Use(middleware.LoggingMiddleware)

	//fetch a project by userId
	apiRoutes.HandleFunc("/projects/{user}", api.HandleGetAllProjectsByUser).Methods("Get")

	//upload file
	apiRoutes.HandleFunc("/uploads/", api.HandleLogFileUpload).Methods("POST")

	apiRoutes.HandleFunc("/uploadSripts/", api.HandleSrciptUpload).Methods("POST")

	//get log file content v2

	apiRoutes.HandleFunc("/v2/content/{fileId}", api.GetLogFileContentv2).Methods("GET")

	//catch the log file updates

	apiRoutes.HandleFunc("/updates", api.HandleFileUpdates).Methods("POST")

	//GetLogsByUserandProject
	apiRoutes.HandleFunc("/logs/getByProject/{id}/", api.GetLogListByProjectID).Methods("GET")

	//Invoke Interpreter
	apiRoutes.HandleFunc("/executeLDEL/{fileId}", api.HandleInvokeELInterpreter).Methods("GET")

	//Get JOSN script result
	apiRoutes.HandleFunc("/executeGetJSON/{fileId}", api.HandleInvokeELInterpreterGetJSON).Methods("GET")
	//Execute LDAL Script
	apiRoutes.HandleFunc("/executeLDAL/{scriptId}", api.HandleExecuteLDAL).Methods("GET")
	//Debug LDAL Queries
	apiRoutes.HandleFunc("/debugLDAL/", api.HandelDebugLDAL).Methods("POST")

	//Craete a project
	apiRoutes.HandleFunc("/project", api.HandleProject).Methods("POST")
	apiRoutes.HandleFunc("/project/{id}", api.GetProjectDetails).Methods("GET")

	//update project
	apiRoutes.HandleFunc("/project/update", api.HandleUpdateProjects).Methods("PUT")

	//delete project
	apiRoutes.HandleFunc("/project/delete/{projectID}", api.HandleDeleteProjects).Methods("DELETE")

	//check project existance
	apiRoutes.HandleFunc("/project/check/{userId}/{projectName}", api.HandleExistProjects).Methods("GET")

	apiRoutes.HandleFunc("/logs/activateLog/{fileId}", api.HandleActiavetLogFile).Methods("GET")

	apiRoutes.HandleFunc("/debug/{projectId}", api.HandelDebugLDEL).Methods("GET")

	apiRoutes.HandleFunc("/debug_save", api.HandleDebugProject).Methods("POST")
	apiRoutes.HandleFunc("/logs/update", api.HandleLogFileUpdate).Methods("PUT")

	//craete a script
	apiRoutes.HandleFunc("/script", api.HandleScripts).Methods("POST")
	//get scripts by projectID
	apiRoutes.HandleFunc("/script/{projectId}", api.HandleGetScriptsByProjectId).Methods("GET")
	apiRoutes.HandleFunc("/getscript/{id}", api.HandleGetScriptDetails).Methods("GET")
	//update scripts
	apiRoutes.HandleFunc("/script/update", api.HandleUpdateScripts).Methods("PUT")

	//craete a json
	apiRoutes.HandleFunc("/customJson", api.HandleCustomJson).Methods("POST")
	//get jsonby projectID
	apiRoutes.HandleFunc("/customJson/{projectId}", api.HandleGetCustomjsonByProjectId).Methods("GET")
	apiRoutes.HandleFunc("/getcustomJson/{id}", api.HandleGetCustomjsonDetails).Methods("GET")
	/*
		*
		 report Routes
	*/
	apiRoutes.HandleFunc("/reports", api.HandleCreateReport).Methods("POST")
	apiRoutes.HandleFunc("/reportQuery", api.HandelReportLDALRequest).Methods("POST")
	apiRoutes.HandleFunc("/reports", api.HandleUpdateReport).Methods("PUT")
	apiRoutes.HandleFunc("/reports/{projectId}", api.HandleGetReportsByProjectId).Methods("GET")
	apiRoutes.HandleFunc("/getReportDetails/{id}", api.HandleGetReportById).Methods("GET")
	apiRoutes.HandleFunc("/reports/{id}", api.HandleDeleteReportById).Methods("DELETE")
	apiRoutes.HandleFunc("/reporttemplate/{id}", api.HandleReportTemplates).Methods("GET")
	apiRoutes.HandleFunc("/reportImage", api.HandleImageupload).Methods("POST")

	return router
}
