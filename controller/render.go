package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderFailed(rw http.ResponseWriter, r *http.Request, errInfo error, dir_failed string) {
	tmpl, err := template.ParseFiles(dir_failed)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = tmpl.Execute(rw, errInfo.Error())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func RenderPage(rw http.ResponseWriter, r *http.Request, data interface{}, dir string) {
	tmpl, err := template.ParseFiles(dir)
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
