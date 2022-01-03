package controller

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/types"
	"log"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
	reportService "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"
)

var reportRepository repository.ReportRepository

func ReportSaveDetails(report datamodels.Report) (interface{}, error) {

	results, err := reportRepository.CreateReport(report)
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("Inserted successgull %v", results.InsertedID)
	return results, err
}

func GetReportById(reportId string) datamodels.Report {
	result := reportRepository.GetReportById(reportId)
	return result
}

func GetReportsByProject(projectId string) []datamodels.Report {

	reports, err := reportRepository.GetAllReportsByProjectId(projectId)
	if err != nil {
		log.Println(err)
	}

	return reports
}

func UpdateReport(updatedReport datamodels.Report) (interface{}, error) {

	result, err := reportRepository.UpdateReport(updatedReport)

	return result, err

}

func DeleteReport(id string) (interface{}, error) {
	result, err := reportRepository.DeleteReportByID(id)
	return result, err
}

func GetReportTemplate(reportId string) types.ReportTemplate {

	reportResult := reportRepository.GetReportById(reportId)
	var reportTemplate types.ReportTemplate
	reportTemplate.ID = reportResult.ID
	reportTemplate.ProjectId = reportResult.ProjectId
	reportTemplate.ReportName = reportResult.ReportName
	reportTemplate.Url = reportResult.Url
	reportTemplate.ReportTemplate = ""

	reportTemplate.ReportTemplate = reportService.GetReportTemplate(reportResult)

	return reportTemplate
}
