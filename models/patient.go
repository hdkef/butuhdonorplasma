package models

type Patient struct {
	ID           string
	Name         string
	Gender       string
	Age          string
	Desc         string
	HospitalName string
	ProvinceID   string
	ProvinceName string
	CityID       string
	CityName     string
	Goldar       string
	Rhesus       string
	Contact      []Contact
}

type Contact struct {
	Name   string
	Tel    string
	Status string
}
