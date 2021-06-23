package find

import (
	"butuhdonorplasma/controller"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var dir string = filepath.Join("public", "find", "find.html")

type FindHandler struct {
}

func GetFindHandler() *FindHandler {
	return &FindHandler{}
}

func (x *FindHandler) Find() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(dir)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data, err := controller.GetProvince()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = tmpl.Execute(rw, data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
