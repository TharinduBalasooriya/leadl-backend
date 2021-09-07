package DB
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client  *mongo.Client
func init() {

	fmt.Println("Database Connection Established")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://tharindu:tharindu@cluster0.vnll5.mongodb.net/myFirstDB?retryWrites=true&w=majority")
	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}
	//project_collection = client.Database("leadldb").Collection(ProjectCollection)

}

func GetDBClient ()*mongo.Client{
	return Client;
	

}
