package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID           primitive.ObjectID `bson:"_id"`
	Date         time.Time          `bson:"date"`
	Name         string             `bson:"name"`
	Gender       string             `bson:"gender"`
	Age          string             `bson:"age"`
	Desc         string             `bson:"desc"`
	HospitalName string             `bson:"hospitalname"`
	ProvinceID   string             `bson:"provinceid"`
	ProvinceName string             `bson:"provincename"`
	CityID       string             `bson:"cityid"`
	CityName     string             `bson:"cityname"`
	Goldar       string             `bson:"goldar"`
	Rhesus       string             `bson:"rhesus"`
	Contact1     Contact            `bson:"cp1"`
	Contact2     Contact            `bson:"cp2"`
}

type Contact struct {
	Name     string `bson:"cpname"`
	Tel      string `bson:"cptel"`
	Relation string `bson:"cprelatioin"`
}
