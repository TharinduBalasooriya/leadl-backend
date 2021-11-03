package repository

import (
	"context"
	"log"
	"time"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	models "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	db "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var COLLECTION_NAME = "reports"
var DATABASE = "leadldb"
var report_collection = new(mongo.Collection)

func init(){
	report_collection = db.Client.Database(Database).Collection(COLLECTION_NAME)
}

type ReportRepository struct{}

func (reportRepository *ReportRepository) CreateReport(report models.Report) (*mongo.InsertOneResult,error){
	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result,err:= report_collection.InsertOne(ctx,report)
	return result,err
}

func (ReportRepository *ReportRepository) GetAllReportsByProjectId(projectId string) ([]models.Report,error){
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	projection := bson.D{{Key:"_id", Value:1}, {Key:"reportname", Value:1}, {Key:"projectid",Value:1}}
	opts := options.Find().SetProjection(projection)
	cursor,err := report_collection.Find(ctx,bson.M{"projectid": projectId},opts)
	var result []models.Report
	if err != nil{
		log.Println(err)
		return result,err
	}

	
	if err = cursor.All(ctx,&result); err != nil{
		log.Println(err)
		return result,err
	}

	return result,err	
}

func (ReportRepository *ReportRepository) GetReportById(reportId string) models.Report{

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(reportId)
	var resultDecode datamodels.Report
	result := report_collection.FindOne(ctx, bson.M{"_id": id})

	result.Decode(&resultDecode)
	return resultDecode

}

func(ReportRepository *ReportRepository) DeleteReportByID(reportId string)(interface{},error){
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	filter := bson.D{{Key:"_id", Value:reportId}}
	result,err := report_collection.DeleteOne(ctx,filter)
	if err != nil{
		log.Println(err)
	}

	return result,err
}


func (ReportRepository *ReportRepository) UpdateReport(report datamodels.Report) (interface{},error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", report.ID}}
	update := bson.D{
		{"$set", bson.D{{"reportname", report.ReportName}}},
		{"$set", bson.D{{"projectid", report.ProjectId}}},
		{"$set", bson.D{{"widgets", report.Widgets}}},
		
	}

	result, err := report_collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		log.Fatal(err)
	}

	

	return result,err

}

