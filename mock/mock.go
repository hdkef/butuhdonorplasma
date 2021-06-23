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
		Contact: []models.Contact{
			{
				Name:   "Name",
				Tel:    "Tel",
				Status: "Status",
			},
			{
				Name:   "Name",
				Tel:    "Tel",
				Status: "Status",
			},
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
