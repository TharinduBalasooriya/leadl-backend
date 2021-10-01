package controller

import (
	"fmt"
	"log"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CustomJsontRepository repository.CustomJsontRepository

func CustomJsonSaveDetails(script datamodels.CustomJson) {

	
	results, err := CustomJsontRepository.CreateCutomJson(script)
	if err != nil {
		log.Fatal(err)

	}

	id := results.(primitive.ObjectID)
	fmt.Println("Successfully inserted" + id.String())

}

func GetCustomJsonDetails(jsonId string) datamodels.CustomJson {
	result := CustomJsontRepository.GetCustomJson(jsonId)
	return result
}

func GetCustomJsontByProject(projectId string) []datamodels.CustomJson {

	customJsons :=CustomJsontRepository.GetCustomJsonsByProjectID(projectId)

	return customJsons
}