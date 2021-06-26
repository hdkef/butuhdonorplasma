package result

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"butuhdonorplasma/konstant"
	"butuhdonorplasma/models"
	"fmt"
	"net/http"
	"path/filepath"
)

var dir string = filepath.Join("public", "result", "result.html")

type ResultHandler struct {
	DBRepo *dbdriver.DBRepo
}

func GetResultHandler(dbrepo *dbdriver.DBRepo) *ResultHandler {
	return &ResultHandler{
		DBRepo: dbrepo,
	}
}

func (x *ResultHandler) Result() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()

		searchKey := models.SearchKey{
			ProvinceID: queryParams.Get(konstant.Provinceid),
			CityID:     queryParams.Get(konstant.Cityid),
		}

		//Find from database
		patients, err := x.DBRepo.FindManyPatients(searchKey)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		controller.RenderPage(rw, r, patients, dir)
	}
}
