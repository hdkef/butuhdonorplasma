package dbdriver

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBUSER string
var DBPASS string
var DBHOST string
var DBPORT string

func init() {
	godotenv.Load()
	DBUSER = os.Getenv("DBUSER")
	DBPASS = os.Getenv("DBPASS")
	DBHOST = os.Getenv("DBHOST")
	DBPORT = os.Getenv("DBPORT")
}

func DBConn(ctx context.Context) (*mongo.Client, error) {

	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/plasma?authSource=admin",
		DBUSER, DBPASS, DBHOST, DBPORT,
	)

	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return client, nil
}
