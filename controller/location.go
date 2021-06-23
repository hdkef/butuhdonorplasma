package controller

import (
	"butuhdonorplasma/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var LocationAPI string

func init() {
	godotenv.Load()
	LocationAPI = os.Getenv("LOCATIONAPI")
}

func GetProvince() ([]models.Province, error) {

	provinceURL := fmt.Sprintf("%s/provinces.json", LocationAPI)

	locationClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, provinceURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	res, err := locationClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var provinces []models.Province

	err = json.NewDecoder(res.Body).Decode(&provinces)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return provinces, nil
}

func GetCity(id string) ([]models.City, error) {

	cityURL := fmt.Sprintf("%s/regencies/%v.json", LocationAPI, id)

	locationClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, cityURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	res, err := locationClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var cities []models.City

	err = json.NewDecoder(res.Body).Decode(&cities)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return cities, nil
}
