package add

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/models"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var dir string = filepath.Join("public", "add", "add.html")
var dir_failed string = filepath.Join("public", "add", "add-failed.html")
var dir_success string = filepath.Join("public", "add", "add-success.html")

type AddHandler struct {
}

func GetAddHandler() *AddHandler {
	return &AddHandler{}
}

func (x *AddHandler) Add() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			handleAddResult(rw, r)
			return
		}

		tmpl, err := template.ParseFiles(dir)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		captcha := controller.GetCaptcha()
		province, err := controller.GetProvince()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var data map[string]interface{} = map[string]interface{}{}
		data["captcha"] = captcha
		data["province"] = province

		err = tmpl.Execute(rw, data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func handleAddResult(rw http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		renderFailed(rw, r)
		return
	}

	patient := models.Patient{
		Name:         r.FormValue("name"),
		Age:          r.FormValue("age"),
		Gender:       r.FormValue("gender"),
		Desc:         r.FormValue("desc"),
		HospitalName: r.FormValue("hospitalname"),
		ProvinceID:   r.FormValue("province"),
		CityID:       r.FormValue("city"),
		Goldar:       r.FormValue("goldar"),
		Rhesus:       r.FormValue("rhesus"),
	}

	fmt.Println(patient)

	renderSuccess(rw, r)
}

func renderSuccess(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(dir_success)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func renderFailed(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(dir_failed)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
