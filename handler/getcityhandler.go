package handler

import (
	"butuhdonorplasma/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCityHandler() http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		var payload struct {
			ID string `json:"id"`
		}

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// data, err := controller.GetCity(payload.ID)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }

		data := []models.City{
			{
				ID:   "1",
				Name: "Bdg",
			},
		}

		err = json.NewEncoder(rw).Encode(&data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
