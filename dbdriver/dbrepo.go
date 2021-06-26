package dbdriver

import (
	"butuhdonorplasma/konstant"
	"butuhdonorplasma/models"
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

func (x *DBRepo) InsertOnePatient(patient models.Patient) (primitive.ObjectID, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//TOBEIMPLEMENTED

	patientbsonD := bson.D{
		{
			Key:   konstant.Provinceid,
			Value: patient.ProvinceID,
		},
		{
			Key:   konstant.Provincename,
			Value: patient.ProvinceName,
		},
		{
			Key:   konstant.Date,
			Value: time.Now().Format("2 Jan 2006 15:04"),
		},
		{
			Key:   konstant.Cityid,
			Value: patient.CityID,
		},
		{
			Key:   konstant.Cityname,
			Value: patient.CityName,
		},
		{
			Key:   konstant.Name,
			Value: patient.Name,
		},
		{
			Key:   konstant.Age,
			Value: patient.Age,
		},
		{
			Key:   konstant.Desc,
			Value: patient.Desc,
		},
		{
			Key:   konstant.Goldar,
			Value: patient.Goldar,
		},
		{
			Key:   konstant.Rhesus,
			Value: patient.Rhesus,
		},
		{
			Key:   konstant.Cpname1,
			Value: patient.Contact1.Name,
		},
		{
			Key:   konstant.Cptel1,
			Value: patient.Contact1.Tel,
		},
		{
			Key:   konstant.Cprelation1,
			Value: patient.Contact1.Tel,
		},
		{
			Key:   konstant.Cpname2,
			Value: patient.Contact2.Name,
		},
		{
			Key:   konstant.Cptel2,
			Value: patient.Contact2.Tel,
		},
		{
			Key:   konstant.Cprelation2,
			Value: patient.Contact2.Tel,
		},
		{
			Key:   konstant.Hospitalname,
			Value: patient.HospitalName,
		},
	}

	res, err := x.DB.Collection(konstant.Col_patients).InsertOne(ctx, patientbsonD)
	if err != nil {
		fmt.Println(err.Error())
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (x *DBRepo) FindManyPatients(searchKey models.SearchKey) ([]models.Patient, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := x.DB.Collection(konstant.Col_patients).Find(ctx, bson.M{
		konstant.Cityid: searchKey.CityID,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var documents []models.Patient

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var document models.Patient
		if err = cursor.Decode(&document); err != nil {
			fmt.Println(err.Error())
			return []models.Patient{}, err
		}
		documents = append(documents, document)
	}

	return documents, nil
}

func (x *DBRepo) DeletePatientByID(id string) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := x.DB.Collection(konstant.Col_patients).DeleteOne(ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return nil, err
	}
	return res.DeletedCount, nil
}
