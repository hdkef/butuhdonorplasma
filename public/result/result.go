package result

import (
	"butuhdonorplasma/dbdriver"
	"butuhdonorplasma/mock"
	"butuhdonorplasma/models"
	"fmt"
	"html/template"
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
			ProvinceID: queryParams.Get("province"),
			CityID:     queryParams.Get("city"),
			Goldar:     queryParams.Get("goldar"),
			Rhesus:     queryParams.Get("rhesus"),
		}

		patients := mock.GetPatientsResult(searchKey)

		tmpl, err := template.ParseFiles(dir)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = tmpl.Execute(rw, patients)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
