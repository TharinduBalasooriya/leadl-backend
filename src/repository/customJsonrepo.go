package repository

import (
	"context"
	"fmt"

	"time"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/db"
)
type CustomJsontRepository struct{}

var customJson_collection = new(mongo.Collection)


const CustomJsonCollection = "CustomJson"


func init(){

	customJson_collection = db.Client.Database(Database).Collection(CustomJsonCollection)

}

func (l *CustomJsontRepository) CreateCutomJson(customJson datamodels.CustomJson) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := customJson_collection.InsertOne(ctx, customJson)
	fmt.Println("\nInserted a JSON: ", result.InsertedID)
	return result.InsertedID, err
}

func (l *CustomJsontRepository) GetCustomJsonsByProjectID(projectId string) []datamodels.CustomJson {

	var Jsons []datamodels.CustomJson

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterCursor, err := customJson_collection.Find(ctx, bson.M{"projectid": projectId})

	if err != nil {
		fmt.Println(err)
	}

	defer filterCursor.Close(ctx)
	for filterCursor.Next(ctx) {

		var json datamodels.CustomJson
		filterCursor.Decode(&json)
		Jsons = append(Jsons, json)
	}

	if err := filterCursor.Err(); err != nil {
		fmt.Println(err.Error())

	}

	return Jsons

}


func (l *CustomJsontRepository) GetCustomJson(jsonId string) datamodels.CustomJson{

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(jsonId)
	var resultDecode datamodels.CustomJson
	result := customJson_collection.FindOne(ctx, bson.M{"_id": id})

	result.Decode(&resultDecode)
	return resultDecode

}
