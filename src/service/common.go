package service

import (
	"log"
	"os")


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