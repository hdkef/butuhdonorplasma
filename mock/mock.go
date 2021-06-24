package mock

import (
	"butuhdonorplasma/models"
	"math/rand"
)

func createOnePatient() models.Patient {
	return models.Patient{
		ID:           string(rand.Int31()),
		Name:         "Name",
		Gender:       "Gender",
		Age:          "Age",
		Desc:         "Desc",
		HospitalName: "HospitalName",
		ProvinceID:   "ProvinceID",
		CityID:       "CityID",
		CityName:     "CityName",
		Goldar:       "Goldar",
		Rhesus:       "Rhesus",
		Contact1: models.Contact{
			Name:     "Name",
			Tel:      "Tel",
			Relation: "Status",
		},
		Contact2: models.Contact{
			Name:     "Name",
			Tel:      "Tel",
			Relation: "Status",
		},
	}
}

func GetPatientsResult(searchKey models.SearchKey) []models.Patient {
	patients := []models.Patient{
		createOnePatient(),
		createOnePatient(),
		createOnePatient(),
		createOnePatient(),
		createOnePatient(),
	}
	return patients
}
