package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/routes"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	//fcllib "github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
)

// LoadEnv /*
func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

}

/*
 Entry point
*/
func main() {

	//Starting the API server
	router := routes.LogRoutes()

	os.MkdirAll("temp", 0755)
	os.MkdirAll("localstorage", 0755)
	os.MkdirAll("debug_env", 0755)

	//Load the env file
	LoadEnv()

	router.Path("/")
	http.Handle("/", router)
	//log.Println("Server Started localhost :3000")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*", "https://leadl-web.herokuapp.com"}))(router)))

	//  result := fcllib.NewFCLWrapper().GetLDALResult("C:\\Users\\thari\\Desktop\\Code2\\tests\\tdpTest\\Defs.txt");
	// log.Println(result)

}
