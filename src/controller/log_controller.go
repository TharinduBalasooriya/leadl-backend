package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"

	//"os"
	"time"

	//"io/ioutil"

	fcllib "github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"
	filestorageHandler "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/filestorage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var logrepo repository.LogRepository
var ldalRepo repository.ScriptRepository

type Loglist struct {
	UserName string   `json:"userName"`
	Project  string   `json:"project"`
	Logs     []string `json:"logs"`
}

type LogContent struct {
	FileName string `json:"filename"`
	Content  string `json:"content"`
}

func GetFileList(user string) []datamodels.Log {

	loglist := logrepo.GetLogsByUser(user)

	return loglist
}

func GetProjects(user string) interface{} {

	projectList := logrepo.GetProjectsByUser(user)

	return projectList
}

func GetLogListByProjectID(projectId string) interface{} {
	logList := logrepo.GetLogsByProject_ID(projectId)
	return logList

}

const (
	S3_REGION = "ap-south-1"
	S3_BUCKET = "leadl"
)

func ExecuteLDEL(fileId string) (interface{}, interface{}) {

	logFileDetails := logrepo.GetLogFileDetails(fileId)
	service.Log_Download_LogFile(fileId)
	service.Log_download_Script(fileId)
	Config_LDEL_DEF(logFileDetails.LogFileName, logFileDetails.FileId)
	service.Log_Execute_LDEL(fileId)
	result := service.Log_Read_Result(fileId)
	JSONresult := service.Log_Read_JSONResult(fileId)

	os.RemoveAll("localstorage/" + fileId)
	return result, JSONresult

}

func Config_LDEL_DEF(logFileName string, fileID string) {

	service.Log_CreateDirectory(fileID)
	service.Log_GetDefFileTempalte(fileID)
	service.Log_Append_LDEL_ScriptLocation(fileID)
	service.Log_Append_LDEL_LogFileLocation(fileID, logFileName)
	service.Log_Append_LDEL_ResultLocation(fileID)
	service.Log_Append_LDEL_JSONResultLocation(fileID)
	service.Log_Append_LDAL_Tree_Location(fileID)
	service.Log_Append_RuleFileLocation(fileID)

}

func GetToActiveDir(fileId string) string {

	logFileDetails := logrepo.GetLogFileDetails(fileId)
	user := logFileDetails.Username
	project := logFileDetails.ProjectId
	var filename = logFileDetails.LogFileName
	var extension = filepath.Ext(filename)
	var logf = filename[0 : len(filename)-len(extension)]

	bucket := "leadl/logs/" + user + "/" + project + "/"

	/*
		TODO:change extension to config
	*/
	item := logf + os.Getenv("BUCKET_ITEM_EXT")
	//item := log + ".txt.zip"

	//fmt.Print(bucket+item)

	object := filestorageHandler.AWS_S3_Object{
		Bucket: bucket,
		Item:   item,
	}

	data := service.Log_GetContent(object, logf, fileId)

	Config_LDEL_DEF(filename, logFileDetails.FileId)

	// Open a new file for writing only
	filePath := "localstorage/" + fileId + "/" + filename
	file, err := os.OpenFile(
		filePath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Println(err)
	}
	//log.Printf("Wrote %d bytes.\n in localstorage", bytesWritten)

	return fileId + " : Activated"

}

func LogUpdateFile(logfile datamodels.Log_Update) string {
	logFileDetails := logrepo.GetLogFileDetails(logfile.FileId)
	user := logFileDetails.Username
	project := logFileDetails.ProjectId
	var filename = logFileDetails.LogFileName

	bucket := "leadl/logs/" + user + "/" + project + "/"
	uploadPath := "logs/" + user + "/" + project + "/" + filename

	item := filename + os.Getenv("ARCHIVED_EXT")

	object := filestorageHandler.AWS_S3_Object{
		Bucket: bucket,
		Item:   item,
	}

	data := service.Log_GetContent(object, filename, logfile.FileId)
	data_string := string(data) + logfile.Content
	newFile := service.ArchiveFile(filename, data_string)

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Println(err)
	}

	/*
	 Create a file storage type object
	*/

	//S3 type object
	s3 := filestorageHandler.AWS_S3{
		Session:   s,
		Filepath:  uploadPath,
		FileBytes: newFile,
	}

	service.Log_uploadFiles(s3)
	return data_string

}

func LogGetFileContentv2(fileId string) interface{} {

	logFileDetails := logrepo.GetLogFileDetails(fileId)
	user := logFileDetails.Username
	project := logFileDetails.ProjectId
	var filename = logFileDetails.LogFileName

	bucket := "leadl/logs/" + user + "/" + project + "/"

	item := filename + os.Getenv("ARCHIVED_EXT")

	object := filestorageHandler.AWS_S3_Object{
		Bucket: bucket,
		Item:   item,
	}

	data := service.Log_GetContent(object, filename, fileId)

	var dataT = string(data)

	logcontent := LogContent{
		FileName: filename,
		Content:  dataT,
	}
	return logcontent
}

func LogSaveDetails(userName string, ProjectId string, logFileName string, fileID string) {

	logfile := datamodels.Log{
		Username:    userName,
		FileId:      fileID,
		LogFileName: logFileName,
		ProjectId:   ProjectId,
		LastUpdate:  time.Now().String(),
	}

	exist, res := logrepo.CheckLogExist(logfile)

	fmt.Println(exist)
	if exist {

		fmt.Println("Log Already Exist")
		logrepo.UpdateTimeStamp(res)

	} else {

		results, err := service.Log_Save_Details(logfile)

		if err != nil {
			log.Println(err)

		}

		id := results.(primitive.ObjectID)
		fmt.Println("Successfully inserted" + id.String())

	}

}

func LogUploadFiles(path string, inputfile multipart.File) {

	// byte array
	fileBytes, err := ioutil.ReadAll(inputfile)
	if err != nil {
		fmt.Println(err)

	}

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Println(err)
	}

	/*
	 Create a file storage type object
	*/

	//S3 type object
	s3 := filestorageHandler.AWS_S3{
		Session:   s,
		Filepath:  path,
		FileBytes: fileBytes,
	}

	service.Log_uploadFiles(s3)

}

type Update struct {
	UserName    string `json:"userName"`
	ProjectName string `json:"project"`
	Data        string `json:"data"`
}

func HandleUpdateData(update Update) {

	fmt.Println(update.UserName)
	fmt.Println(update.ProjectName)
	fmt.Println(update.Data)

}

func ExecuteLDAL(scriptId string) string {
	

	var ldalDetails datamodels.LDALscript
	ldalDetails = ldalRepo.GetLDALScripts(scriptId)

	logFileDetails := logrepo.GetLogFileDetails(ldalDetails.BoundedId)
	service.Log_Download_LogFile(ldalDetails.BoundedId)
	service.Log_download_Script(ldalDetails.BoundedId)
	Config_LDEL_DEF(logFileDetails.LogFileName, logFileDetails.FileId)
	service.Log_Execute_LDEL(ldalDetails.BoundedId)
	decodedContent, err := base64.StdEncoding.DecodeString(ldalDetails.Content)
	if err != nil {
		log.Println("decode error:", err)

	}
	service.WriteToFile("localstorage/"+logFileDetails.FileId, "LDAL_Script.txt", string(decodedContent))
	result := fcllib.NewFCLWrapper().GetLDALResult("localstorage/" + ldalDetails.BoundedId + "/" + "Defs.txt")
	
	os.RemoveAll("localstorage/" + logFileDetails.FileId)

	return result

}
