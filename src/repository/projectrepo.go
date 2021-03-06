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

var project_collection = new(mongo.Collection)

const ProjectCollection = "Project"
const Database = "leadldb"

func init(){

	project_collection = db.Client.Database(Database).Collection(ProjectCollection)

}


type ProjectRepository struct{}

func (l *ProjectRepository) SaveProject(project datamodels.Project) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := project_collection.InsertOne(ctx, project)
	fmt.Println("\nInserted a single project: ", result.InsertedID)
	return result.InsertedID, err
}

func (l *ProjectRepository) CheckprojectExist(project datamodels.Project) (bool, string) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := project_collection.FindOne(ctx, bson.M{"_id": project.ProjectId})

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

func (l *ProjectRepository) GetProjectsByUserV2(userId string) []datamodels.Project {

	var projects []datamodels.Project

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterCursor, err := project_collection.Find(ctx, bson.M{"userid": userId})

	if err != nil {
		fmt.Println(err)
	}

	defer filterCursor.Close(ctx)
	for filterCursor.Next(ctx) {

		var project datamodels.Project
		filterCursor.Decode(&project)
		projects = append(projects, project)
	}

	if err := filterCursor.Err(); err != nil {
		fmt.Println(err.Error())

	}

	return projects

}

func (l *ProjectRepository) UpadteProject(project datamodels.Project) interface{} {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", project.ProjectId}}
	update := bson.D{
		{"$set", bson.D{{"projectname", project.ProjectName}}},
		{"$set", bson.D{{"userid", project.UserId}}},
		{"$set", bson.D{{"expiredate", project.ExpireDate}}},
		{"$set", bson.D{{"scriptstatus", project.ScriptStatus}}},
		{"$set", bson.D{{"script", project.Script}}},
	}

	result, err := project_collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	return result

}

func (l *ProjectRepository) DeleteProject(projectId string) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fmt.Print(projectId)
	id, _ := primitive.ObjectIDFromHex(projectId)
	result, err := project_collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	return result

}

func (l *ProjectRepository) CheckprojectExistByUser(projectName string , userId string) interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result := project_collection.FindOne(ctx, bson.M{"userid": userId,"projectname": projectName})

	var resultLog bson.M
	result.Decode(&resultLog)

	// ErrNoDocuments means that the filter did not match any documents in the collection
	if resultLog == nil {
		fmt.Println("\n project not exists")
		return false

	} else {
		fmt.Println("\n project exists")
		return true

	}

}


func (l *ProjectRepository) GetProjectDetails(projectId string) datamodels.Project {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	id, _ := primitive.ObjectIDFromHex(projectId)
	var resultDecode datamodels.Project
	result := project_collection.FindOne(ctx, bson.M{"_id": id})

	result.Decode(&resultDecode)
	return resultDecode

}