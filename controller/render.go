package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderFailed(rw http.ResponseWriter, r *http.Request, errInfo error, tmpl *template.Template) {

	err := tmpl.Execute(rw, errInfo.Error())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func RenderPage(rw http.ResponseWriter, r *http.Request, data interface{}, tmpl *template.Template) {

	err := tmpl.Execute(rw, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
