package result

import (
	"butuhdonorplasma/mock"
	"butuhdonorplasma/models"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var dir string = filepath.Join("public", "result", "result.html")

type ResultHandler struct {
}

func GetResultHandler() *ResultHandler {
	return &ResultHandler{}
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
