package api

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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
	result, err := controller.ReportSaveDetails(report)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
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
	result, err := controller.UpdateReport(report)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}

}

func HandleDeleteReportById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := controller.DeleteReport(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(result)
	}

}

func HandleReportTemplates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := controller.GetReportTemplate(params["id"])
	json.NewEncoder(w).Encode(result)
}

func HandleImageupload(w http.ResponseWriter, r *http.Request) {
	maxSize := int64(1024000) // allow only 1MB of file size

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Image too large. Max Size: %v", maxSize)
		return
	}

	file, fileHeader, err := r.FormFile("report-img")
	if err != nil {
		log.Println(err)
		//
		fmt.Fprintf(w, "Could not get uploaded file")
		return
	}
	defer file.Close()

	// create an AWS session which can be
	// reused if we're uploading many files
	s, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)
	if err != nil {
		//
		fmt.Fprintf(w, "Could not upload file")
	}

	fileName, err := service.UploadFileToS3(s, file, fileHeader, "leadl-images")

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		type fileUploadResponse struct {
		 
			Url string  `json:"url"`
			}
			
			respone:=&fileUploadResponse{

				Url: "https://leadl-images.s3.ap-south-1.amazonaws.com/"+fileName,
			}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respone)
	}

}