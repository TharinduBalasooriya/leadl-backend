package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"bytes"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/globalsign/mgo/bson"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func WriteToFile(location string, fileName string, content string) {
	os.MkdirAll(location, 0755)
	file, err := os.OpenFile(
		location+"/"+fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	file.WriteString(content)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

}

func ReadFromFile(location string) string {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}



// UploadFileToS3 saves a file to aws bucket and returns the url to the file and an error if there's any
func UploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader  ,s3bucket string)  (string, error) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	// create a unique file name for the file
	tempFileName :=  bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)
	
	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s3bucket),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"),// could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		// ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	
	})
	if err != nil {
		return "", err 
	}

	return tempFileName, err
}


