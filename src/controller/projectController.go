package controller

import (
	"fmt"
	"log"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var projectrepo repository.ProjectRepository

func ProjectSaveDetails(project datamodels.Project) {

	exist, res := projectrepo.CheckprojectExist(project)

	if exist {

		fmt.Println("Project Already Exist")
		fmt.Print(res)

	} else {

		results, err := service.Project_Save_Details(project)

		if err != nil {
			log.Println(err)

		}

		id := results.(primitive.ObjectID)
		fmt.Println("Successfully inserted" + id.String())

	}

}

func UpdateProject(project datamodels.Project) interface{} {

	results := projectrepo.UpadteProject(project)

	return results
}

func GetProjectsV2(user string) []datamodels.Project {
	projectList := projectrepo.GetProjectsByUserV2(user)
	return projectList
}

func DeleteProject(projectId string) interface{} {
	results := projectrepo.DeleteProject(projectId)
	return results
}

func CheckProject(projectId string, userId string) interface{} {
	results := projectrepo.CheckprojectExistByUser(userId, projectId)
	return results
}

func GetProjectDetails(projectId string) datamodels.Project {
	result := projectrepo.GetProjectDetails(projectId)
	return result
}
