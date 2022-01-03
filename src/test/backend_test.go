package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	models "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	repository "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)




func TestCreateDir(t *testing.T) {
    service.Log_CreateDirectory("12212");

}

func TestCreateReport(t *testing.T){

    var reportRepo  repository.ReportRepository
    var newReport models.Report
    var newWidget models.Widget
    newReport.ReportName="ReportTest"
    newReport.ProjectId="project123"
    newWidget.ScriptId="TestScript1"
    newWidget.Type="pie"
    newWidget.X="123"
    newWidget.Y="23"
    newReport.Headers= []models.Header{{Name: "accessToke",JsFunction: "function () {retun true} "}, {Name: "cookie",JsFunction: "function () {retun true} "}}

    newReport.Widgets = append(newReport.Widgets,newWidget)
    result,err := reportRepo.CreateReport(newReport)
     if(result == nil || err != nil){
        t.Fatalf(`Report create test failed failed %v`,err)
     }
}

func TestGetAllReportByProjectId(t *testing.T){
    var reportRepo  repository.ReportRepository
    projectId := "project123"
    var reports []models.Report
    reports,err := reportRepo.GetAllReportsByProjectId(projectId)
    if(err != nil || reports == nil){
        t.Fatalf(`Get reports by project id failed due to %v`,err)
    }
    result, err := json.Marshal(reports)
    if(err != nil){
        log.Println(err)
    }
    log.Println(string(result))


}

func TestDeleteReportById(t *testing.T){

    var reportRepo  repository.ReportRepository
    var newReport models.Report
    var newWidget models.Widget
    newReport.ReportName="ReportTestDelete"
    newReport.ProjectId="project123"
    newWidget.ScriptId="TestScript1"
    newWidget.Type="pie"
    newWidget.X="123"
    newWidget.Y="23"

    newReport.Widgets = append(newReport.Widgets,newWidget)
    var result *mongo.InsertOneResult
    result,err := reportRepo.CreateReport(newReport)
     if(result == nil || err != nil){
        t.Fatalf(`Report delete test failed failed %v`,err)
     }

    var resultId interface{} = result.InsertedID
    strId := fmt.Sprintf("%v",resultId)
     delResult,err := reportRepo.DeleteReportByID(strId)
     if err != nil{
        t.Fatalf(`Report delete test failed failed %v`,err)
     }
     fmt.Println(delResult)


}

func TestUpdateDocument(t *testing.T){
    var reportRepo  repository.ReportRepository
    var newReport models.Report
    var newWidget models.Widget
    newReport.ReportName="ReportTestUpdate"
    newReport.ProjectId="project123"
    newWidget.ScriptId="TestScript1"
    newWidget.Type="pie"
    newWidget.X="123"
    newWidget.Y="23"

    newReport.Widgets = append(newReport.Widgets,newWidget)
    var result *mongo.InsertOneResult
    result,err := reportRepo.CreateReport(newReport)
     if(result == nil || err != nil){
        t.Fatalf(`Report delete test failed failed %v`,err)
     }

    var resultId interface{} = result.InsertedID
    strId := fmt.Sprintf("%v",resultId)
    newReport.ID,err = primitive.ObjectIDFromHex(strId)
    if err != nil{
        log.Println(err)
    }
    //Add a another widget
    newReport.Widgets = append(newReport.Widgets,newWidget)
    newReport.ReportName ="NewName"
    updatedResult,err := reportRepo.UpdateReport(newReport)
    if err != nil{
        log.Printf("Update test failed %v",updatedResult)
    }

    stringRes, err := json.Marshal(updatedResult)
    if(err != nil){
        log.Println(err)
    }
    log.Println(string(stringRes))


}