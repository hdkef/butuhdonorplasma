package controller

import (
	"butuhdonorplasma/models"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/jszwec/csvutil"
)

var provinces []models.Province
var PROVINCES = filepath.Join("data", "provinces.csv")
var REGENCIES = filepath.Join("data", "regencies.csv")

func init() {
	godotenv.Load()
	provinces = retrieveProvinceOnce()
}

func retrieveProvinceOnce() []models.Province {
	provinces, err := decodeProvinces()
	if err != nil {
		panic(err.Error())
	}
	return provinces
}

func GetProvince() []models.Province {
	return provinces
}

func GetCity(id string) ([]models.City, error) {
	return decodeCitiesByProvinceID(id)
}

func readFile(filepath string) ([]byte, error) {

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return file, nil
}

func decodeProvinces() ([]models.Province, error) {

	csvInput, err := readFile(PROVINCES)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var provinces []models.Province

	if err := csvutil.Unmarshal(csvInput, &provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}

func openFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func decodeCitiesByProvinceID(ID string) ([]models.City, error) {

	file, err := openFile(REGENCIES)
	if err != nil {
		fmt.Println(err.Error())
		return []models.City{}, err
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		fmt.Println(err.Error())
		return []models.City{}, err
	}

	var cities []models.City

	for {
		var city models.City
		err = dec.Decode(&city)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if city.ProvinceID == ID {
			cities = append(cities, city)
		}
	}

	if len(cities) == 0 {
		return []models.City{}, errors.New("NO CITIES FOUND")
	}
	return cities, nil
}
