package dbdriver

import (
	"butuhdonorplasma/konstant"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBUSER string
var DBPASS string
var DBHOST string
var DBPORT string
var DBNAME string

func init() {
	godotenv.Load()
	DBUSER = os.Getenv("DBUSER")
	DBPASS = os.Getenv("DBPASS")
	DBHOST = os.Getenv("DBHOST")
	DBPORT = os.Getenv("DBPORT")
	DBNAME = os.Getenv("DBNAME")
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

	err = CreateIndex(client) //creating index after connecting
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateIndex(client *mongo.Client) error {

	ctx, _ := context.WithTimeout(context.Background(), 25*time.Second)

	mod := mongo.IndexModel{
		Keys: bson.M{
			konstant.Cityid: 1, //create index for city id for fast query
		}, Options: options.Index().SetExpireAfterSeconds(1209600), //Time to live for automatically delete after 1 week
	}

	_, err := client.Database(DBNAME).Collection(konstant.Col_patients).Indexes().CreateOne(
		ctx,
		mod,
	)
	if err != nil {
		fmt.Println("error index", err)
		panic(err.Error())
	}
	return nil
}
