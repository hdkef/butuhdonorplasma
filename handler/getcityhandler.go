package handler

import (
	"butuhdonorplasma/controller"
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

		data, err := controller.GetCity(payload.ID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = json.NewEncoder(rw).Encode(&data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
