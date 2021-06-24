package models

import "time"

type Patient struct {
	ID           string
	Date         time.Time
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
	Contact1     Contact
	Contact2     Contact
}

type Contact struct {
	Name     string
	Tel      string
	Relation string
}
