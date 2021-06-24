package find

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var dir string = filepath.Join("public", "find", "find.html")

type FindHandler struct {
	DBRepo *dbdriver.DBRepo
}

func GetFindHandler(dbrepo *dbdriver.DBRepo) *FindHandler {
	return &FindHandler{
		DBRepo: dbrepo,
	}
}

func (x *FindHandler) Find() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(dir)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := controller.GetProvince()

		err = tmpl.Execute(rw, data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
