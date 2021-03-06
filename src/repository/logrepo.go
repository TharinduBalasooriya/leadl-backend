package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"time"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/db"
)

var log_collection = new(mongo.Collection)

const LogsCollection = "Logs"

/*
	Initalizing database configeration
*/

func init() {

	log_collection = db.Client.Database(Database).Collection(LogsCollection)

}



type LogRepository struct{}

func (l *LogRepository) SaveLog(log datamodels.Log) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := log_collection.InsertOne(ctx, log)
	fmt.Println("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err

}

func (l *LogRepository) CheckLogExist(logfile datamodels.Log) (bool, string) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result := log_collection.FindOne(ctx, bson.M{"username": logfile.Username, "projectId": logfile.ProjectId, "logfilename": logfile.LogFileName})

	var resultLog bson.M

	result.Decode(&resultLog)

	/*
		check existences
	*/
	if len(resultLog) == 0 {

		return false, ""

	} else {
		stringObjectId := resultLog["_id"].(primitive.ObjectID).Hex()
		return true, stringObjectId
	}

}

func (l *LogRepository) UpdateTimeStamp(objectId string) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	fmt.Println(objectId)
	id, _ := primitive.ObjectIDFromHex(objectId)

	result, err := log_collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"lastupdate", time.Now().String()}}},
		},
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Time stamp updated %v", result.MatchedCount)

}

/*
*
*	TODO : Change to work with userId instead
	of user Name
*
*/

func (l *LogRepository) GetLogsByUser(username string) []datamodels.Log {

	var logs []datamodels.Log

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterCursor, err := log_collection.Find(ctx, bson.M{"username": username})

	if err != nil {
		fmt.Println(err)
	}

	defer filterCursor.Close(ctx)
	for filterCursor.Next(ctx) {

		var log datamodels.Log
		filterCursor.Decode(&log)
		logs = append(logs, log)
	}

	if err := filterCursor.Err(); err != nil {
		fmt.Println(err.Error())

	}

	return logs

}

func (l *LogRepository) GetLogsByProject_ID(projectID string) []datamodels.Log {

	var logs []datamodels.Log

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterCursor, err := log_collection.Find(ctx, bson.M{"projectid": projectID})

	if err != nil {
		fmt.Println(err)
	}

	defer filterCursor.Close(ctx)
	for filterCursor.Next(ctx) {
		var log datamodels.Log
		filterCursor.Decode(&log)
		logs = append(logs, log)
	}

	if err := filterCursor.Err(); err != nil {
		fmt.Println(err.Error())

	}

	return logs

}

func (l *LogRepository) GetProjectsByUser(username string) interface{} {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	projetcs, err := log_collection.Distinct(ctx, "projectname", bson.D{{"username", username}})

	if err != nil {
		fmt.Println(err)
	}

	return projetcs

}

func (l *LogRepository) GetLogFileDetails(fileId string) datamodels.Log {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var resultDecode datamodels.Log
	result := log_collection.FindOne(ctx, bson.M{"fileid": fileId})

	result.Decode(&resultDecode)
	return resultDecode

}
