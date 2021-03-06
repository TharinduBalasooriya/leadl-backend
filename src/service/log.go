package service

import (

	//Importing file storage utility
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	fclLib "github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
	filestorageHandler "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/filestorage"
)

/*
This package containes all business logic log file

*/

var logrepo repository.LogRepository
var cusjsonrepo repository.CustomJsontRepository

func unzipLogfile(logFIleName string, fileId string) {

	//fmt.Println("temp/" + logFIleName + os.Getenv("BUCKET_ITEM_EXT"))

	zipReader, err := zip.OpenReader("temp/" + fileId + "/" + logFIleName + os.Getenv("ARCHIVED_EXT"))
	if err != nil {
		log.Println(err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			log.Println(err)
		}
		defer zippedFile.Close()

		targetDir := "./temp/" + fileId
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since not a directory
			log.Println("Extracting file:", file.Name)

			// Open an output file for writing
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Println(err)
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func unzipLogfilev2(logfilename string) {

	fmt.Println("temp/" + logfilename + os.Getenv("BUCKET_ITEM_EXT"))

	zipReader, err := zip.OpenReader("temp/" + logfilename + os.Getenv("BUCKET_ITEM_EXT"))
	if err != nil {
		log.Println(err)
	}
	defer zipReader.Close()

	// Iterate through each file/dir found in
	for _, file := range zipReader.Reader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Println(err)
		}
		defer zippedFile.Close()

		// Specify what the extracted file name should be.
		// You can specify a full path or a prefix
		// to move it to a different directory.
		// In this case, we will extract the file from
		// the zip to a file of the same name.
		targetDir := "./temp"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since not a directory
			log.Println("Extracting file:", file.Name)

			// Open an output file for writing
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Println(err)
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

/*
	Upload a file
*/
func Log_uploadFiles(fs filestorageHandler.FileStorage) {

	err := fs.AddFiles() // calling add files function of the file storage
	if err != nil {
		log.Println(err)
	}

}

//Save Log Details in mongo db
func Log_Save_Details(log datamodels.Log) (interface{}, error) {

	resultID, err := logrepo.SaveLog(log)
	return resultID, err

}

func Log_GetContent(file_object filestorageHandler.File, logfileName string, fileId string) []byte {

	//fileExtension := os.Getenv("FILE_EXT")
	//fileExtension := ".txt"

	err := file_object.GetContent(fileId)
	if err != nil {
		log.Println(err)
	}
	unzipLogfile(logfileName, fileId)

	data, err := ioutil.ReadFile("temp/" + fileId + "/" + logfileName)
	if err != nil {
		panic(err)
	}
	os.RemoveAll("temp/" + fileId)

	return data

}

func Log_GetContentV2(file_object filestorageHandler.File, logfileName string, fileId string) []byte {

	err := file_object.GetContent(fileId)
	if err != nil {
		log.Println(err)
	}
	unzipLogfilev2(logfileName)

	data, err := ioutil.ReadFile("temp/" + logfileName)
	if err != nil {
		panic(err)
	}

	return data

}

//Create local storage derectories

func Log_CreateDirectory(fileId string) {

	path := "localstorage/" + fileId
	err := os.MkdirAll(path, 0755)

	if err != nil {
		log.Println(err)
	}
}

func Log_GetDefFileTempalte(fileId string) {

	//Open DefFile template

	defFileTemplate, err := os.Open("util/templates/Defs.txt")
	if err != nil {
		log.Println(err)
	}

	defer defFileTemplate.Close()

	//Create New File

	newFilePath := "localstorage/" + fileId + "/Defs.txt"
	newFile, err := os.Create(newFilePath)
	if err != nil {
		log.Println(err)
	}
	defer newFile.Close()

	//Copy bytes create a new Template

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, defFileTemplate)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Println(err)
	}

}

func Log_Execute_LDEL(fileId string) {

	defFilePath := "localstorage/" + fileId + "/Defs.txt"

	fclLib.NewFCLWrapper().RunELInterpretter(defFilePath)
}

func Log_Download_LogFile(fileId string,requestId string) {
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

	data := Log_GetContent(object, filename, fileId)
	os.MkdirAll("localstorage/"+requestId, 0755)
	file, err := os.OpenFile(
		"localstorage/"+requestId+"/"+filename,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	bytesWritten, err := file.Write(data)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

}

func Log_download_Script(fileId string,requestId string) {
	logFileDetails := logrepo.GetLogFileDetails(fileId)
	projectDetails := projectrepo.GetProjectDetails(logFileDetails.ProjectId)
	script, err := base64.StdEncoding.DecodeString(projectDetails.Script)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	os.MkdirAll("localstorage/"+requestId, 0755)
	file, err := os.OpenFile(
		"localstorage/"+requestId+"/script.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	bytesWritten, err := file.Write(script)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

}

func DownloadCustomJSON(boundedID string,requestId string){
	fmt.Println(boundedID)
	customJSONRequest := cusjsonrepo.GetCustomJson(boundedID)
	
	

	os.MkdirAll("localstorage/"+requestId, 0755)
	file, err := os.OpenFile(
		"localstorage/"+requestId+"/result.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	bytesWritten, err := file.Write([]byte(customJSONRequest.Content))
	if err != nil {
		log.Println(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

}



func Log_Read_Result(fileId string) interface{} {
	resultFilePath := "localstorage/" + fileId + "/result.txt"

	// Open file for reading
	file, err := os.Open(resultFilePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	type Response struct {
		FileId string `json:"fileId"`
		Result string `json:"result"`
	}

	response := Response{}

	response.FileId = fileId
	response.Result = string(data)

	return response

}

func Log_Read_JSONResult(fileId string) interface{} {
	resultFilePath := "localstorage/" + fileId + "/JSONresult.json"

	// Open file for reading
	file, err := os.Open(resultFilePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	type Response struct {
		FileId string `json:"fileId"`
		Result string `json:"result"`
	}

	response := Response{}

	response.FileId = fileId
	response.Result = string(data)

	return response

}

func Log_Append_LDEL_ScriptLocation(fileId string) {

	defFileLocation := "localstorage/" + fileId + "/Defs.txt"
	newDef := "DEF	LDEL_SCRIPT_FILE			../src/localstorage/" + fileId + "/script.txt\n"

	defFile, err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer defFile.Close()

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}

}

func Log_Append_LDEL_LogFileLocation(fileId string, fileName string) {

	defFileLocation := "localstorage/" + fileId + "/Defs.txt"
	newDef := "DEF	LDEL_LOG_FILE				../src/localstorage/" + fileId + "/" + fileName + "\n"

	defFile, err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer defFile.Close()

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}

}

func Log_Append_LDEL_ResultLocation(fileId string) {

	defFileLocation := "localstorage/" + fileId + "/Defs.txt"
	newDef := "DEF	LDEL_RESULT_FILE			../src/localstorage/" + fileId + "/result.txt\n"

	defFile, err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer defFile.Close()

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}

}

func Log_Append_LDEL_JSONResultLocation(fileId string) {

	defFileLocation := "localstorage/" + fileId + "/Defs.txt"
	newDef := "DEF	LDEL_RESULT_JSONFILE			../src/localstorage/" + fileId + "/JSONresult.json\n"

	defFile, err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer defFile.Close()

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}

}

func Log_Append_LDAL_Tree_Location(fileId string){


	defFileLocation := "localstorage/" + fileId + "/Defs.txt"
	newDef := "DEF	 TREE_LOCATION 			../src/localstorage/" + fileId + "/result.txt\n"

	defFile, err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer defFile.Close()

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}

}
func Log_Append_DebugJSONLocation(requestID string){

	defFileLocation := "localstorage/" + requestID + "/Defs.txt"
	newDef := "DEF	 DEBUG_JSON 			../src/localstorage/" + requestID + "/Debug_Result.json\n"
	

	defFile, err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer defFile.Close()

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}
	
}
func Log_Append_RuleFileLocation(fileId string){


	defFileLocation := "localstorage/" + fileId + "/Defs.txt"
	newDef := "DEF	 RULE_FILE_NAME			../src/localstorage/" + fileId + "/LDAL_Script.txt\n"
	

	defFile, err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer defFile.Close()

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}

}

func ArchiveFile(fileName string, content string) []byte {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	f, err := w.Create(fileName)

	if err != nil {
		log.Println(err)
	}
	_, err = f.Write([]byte(content))
	if err != nil {
		log.Println(err)
	}
	err = w.Close()
	if err != nil {
		log.Println(err)
	}

	val := buf.Bytes()

	return val

}
