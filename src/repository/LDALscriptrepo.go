package repository

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	db "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/db"
)

var script_collection = new(mongo.Collection)

const ScriptCollection = "Script"

func init(){

	script_collection = db.Client.Database(Database).Collection(ScriptCollection)

}


// func init() {

// 	fmt.Println("Database Connection Established")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	clientOptions := options.Client().ApplyURI("mongodb+srv://tharindu:tharindu@cluster0.vnll5.mongodb.net/myFirstDB?retryWrites=true&w=majority")
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	script_collection = client.Database("leadldb").Collection(ScriptCollection)

// }

type ScriptRepository struct{}

func (l *ScriptRepository) SaveLDALScript(LDALscript datamodels.LDALscript) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := script_collection.InsertOne(ctx, LDALscript)
	fmt.Println("\nInserted a single LDALscript: ", result)
	return result, err
}

func (l *ScriptRepository) CheckScriptExistance(script datamodels.LDALscript) (bool, string) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := script_collection.FindOne(ctx, bson.M{"_id": script.SciptId})

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

func (l *ScriptRepository) GetScriptsByProjectID(projectId string) []datamodels.LDALscript {

	var scripts []datamodels.LDALscript

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterCursor, err := script_collection.Find(ctx, bson.M{"projectid": projectId})

	if err != nil {
		fmt.Println(err)
	}

	defer filterCursor.Close(ctx)
	for filterCursor.Next(ctx) {

		var script datamodels.LDALscript
		filterCursor.Decode(&script)
		scripts = append(scripts, script)
	}

	if err := filterCursor.Err(); err != nil {
		fmt.Println(err.Error())

	}

	return scripts

}

func (l *ScriptRepository) UpadteLDALScript(script datamodels.LDALscript) interface{} {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", script.SciptId}}
	update := bson.D{
		{"$set", bson.D{{"scriptname", script.ScriptName}}},
		{"$set", bson.D{{"projectid", script.ProjectID}}},
		{"$set", bson.D{{"boundstatus", script.BoundStatus}}},
		{"$set", bson.D{{"boundedid", script.BoundedId}}},
		{"$set", bson.D{{"content", script.Content}}},
		{"$set",bson.D{{"logquery",script.LogQuery}}},
	}

	result, err := script_collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	return result

}

func (l *ScriptRepository) GetLDALScripts(scriptId string) datamodels.LDALscript {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(scriptId)
	var resultDecode datamodels.LDALscript
	result := script_collection.FindOne(ctx, bson.M{"_id": id})

	result.Decode(&resultDecode)
	return resultDecode

}
