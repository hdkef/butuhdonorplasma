package delete

import (
	"butuhdonorplasma/dbdriver"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type DeleteHandler struct {
	DBRepo *dbdriver.DBRepo
}

var dir string = filepath.Join("public", "delete", "delete.html")

func GetDeleteHandler(dbrepo *dbdriver.DBRepo) *DeleteHandler {
	return &DeleteHandler{
		DBRepo: dbrepo,
	}
}

func (x *DeleteHandler) Delete() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(dir)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = tmpl.Execute(rw, 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
