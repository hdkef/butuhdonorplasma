package index

import (
	"butuhdonorplasma/dbdriver"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var dir string = filepath.Join("public", "index", "index.html")

type IndexHandler struct {
	DBRepo *dbdriver.DBRepo
}

func GetIndexHandler(dbrepo *dbdriver.DBRepo) *IndexHandler {
	return &IndexHandler{
		DBRepo: dbrepo,
	}
}

func (x *IndexHandler) Index() http.HandlerFunc {
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
