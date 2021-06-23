package result

import (
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

		//TOBEIMPLEMENTED
		//get queryparams

		fmt.Println(queryParams)

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
