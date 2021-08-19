package routes

import (
	"fmt"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/websocket"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/api"
	"github.com/gorilla/mux"
)

func LogRoutes() (*mux.Router ){

	router := mux.NewRouter().StrictSlash(true)
	apiRoutes := router.PathPrefix("/api").Subrouter().StrictSlash(true)



	//Get All Log files
	router.HandleFunc("/ws", websocket.WSPage).Methods("GET")
	router.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw,"HomeRoute")
	})

	//apiRoutes.Use(middleware.LoggingMiddleware)
	apiRoutes.HandleFunc("/logs/{user}/", api.GetAllLog).Methods("GET")

	//getAllProjetcs
	apiRoutes.HandleFunc("/projects/{user}/", api.GetAllProjects).Methods("GET")

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



	//Craete a project
	apiRoutes.HandleFunc("/project", api.HandleProject).Methods("POST")
	apiRoutes.HandleFunc("/project/{id}", api.GetProjectDetails).Methods("GET")
	//fetch a project by userId
	apiRoutes.HandleFunc("/projectV2/{user}", api.GetAllProjectsV2).Methods("Get")

	//update project
	apiRoutes.HandleFunc("/project/update", api.HandleUpdateProjects).Methods("PUT")

	//delete project
	apiRoutes.HandleFunc("/project/delete/{projectID}", api.HandleDeleteProjects).Methods("DELETE")

	//check project existance 
	apiRoutes.HandleFunc("/project/check/{userId}/{projectName}", api.HandleExistProjects).Methods("GET")

	apiRoutes.HandleFunc("/logs/activateLog/{fileId}", api.HandleActiavetLogFile).Methods("GET")

	apiRoutes.HandleFunc("/debug/{projectId}", api.HandelDebugLDEL).Methods("GET")

	apiRoutes.HandleFunc("/debug_save", api.HandleDebugProject).Methods("POST")
	apiRoutes.HandleFunc("/logs/update",api.HandleLogFileUpdate).Methods("PUT")
	

	return router
}