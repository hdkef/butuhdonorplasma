package dbdriver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBRepo struct {
	DB *mongo.Database
}

func GetDBRepo(db *mongo.Database) *DBRepo {
	return &DBRepo{
		DB: db,
	}
}

func (x *DBRepo) InsertOne(col string, data bson.D) (string, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res, err := x.DB.Collection(col).InsertOne(ctx, data)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func (x *DBRepo) FindMany(col string, filter bson.M) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := x.DB.Collection(col).Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var document bson.M
		if err = cursor.Decode(&document); err != nil {
			fmt.Println(err.Error())
			return "", err
		}
		fmt.Println(document)
	}

	return nil, nil
}

func (x *DBRepo) Delete(col string) error {
	return nil
}
