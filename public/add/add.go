package add

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"butuhdonorplasma/konstant"
	"butuhdonorplasma/models"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

var dir string = filepath.Join("public", "add", "add.html")
var dir_failed string = filepath.Join("public", "utils", "fail.html")
var dir_success string = filepath.Join("public", "add", "add-success.html")

type AddHandler struct {
	DBRepo *dbdriver.DBRepo
}

func GetAddHandler(dbrepo *dbdriver.DBRepo) *AddHandler {
	return &AddHandler{
		DBRepo: dbrepo,
	}
}

func (x *AddHandler) Add() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			handleAddPost(rw, r, x)
			return
		}

		captcha := controller.GetCaptcha()
		provinces := controller.GetProvince()

		var data map[string]interface{} = map[string]interface{}{}
		data["captcha"] = captcha
		data["province"] = provinces

		controller.RenderPage(rw, r, data, dir)
	}
}

func handleAddPost(rw http.ResponseWriter, r *http.Request, x *AddHandler) {

	err := r.ParseForm()
	if err != nil {
		controller.RenderFailed(rw, r, err, dir_failed)
		return
	}

	captchaID, err := strconv.Atoi(r.FormValue("captchaid"))
	if err != nil {
		controller.RenderFailed(rw, r, err, dir_failed)
		return
	}
	captcha := r.FormValue("captcha")

	err = controller.CheckCaptcha(int64(captchaID), captcha)
	if err != nil {
		controller.RenderFailed(rw, r, err, dir_failed)
		return
	}

	patient := models.Patient{
		Name:         r.FormValue(konstant.Name),
		Age:          r.FormValue(konstant.Age),
		Desc:         r.FormValue(konstant.Desc),
		HospitalName: r.FormValue(konstant.Hospitalname),
		ProvinceID:   r.FormValue(konstant.Provinceid),
		ProvinceName: r.FormValue(konstant.Provincename),
		CityID:       r.FormValue(konstant.Cityid),
		CityName:     r.FormValue(konstant.Cityname),
		Goldar:       r.FormValue(konstant.Goldar),
		Rhesus:       r.FormValue(konstant.Rhesus),
		Contact1: models.Contact{
			Name:     r.FormValue(konstant.Cpname1),
			Tel:      r.FormValue(konstant.Cptel1),
			Relation: r.FormValue(konstant.Cprelation1),
		},
	}

	cpname2 := r.FormValue(konstant.Cpname2)

	fmt.Println(patient)

	if cpname2 != "" {
		patient.Contact2.Name = cpname2
		patient.Contact2.Tel = r.FormValue(konstant.Cptel2)
		patient.Contact2.Tel = r.FormValue(konstant.Cprelation2)
	}

	if patient.Name == "" || patient.Age == "" || patient.Desc == "" || patient.HospitalName == "" || patient.ProvinceID == "" || patient.CityID == "" || patient.Goldar == "" || patient.Rhesus == "" {
		controller.RenderFailed(rw, r, errors.New("something is missing"), dir_failed)
		return
	}

	//ADD TO DATABASE
	id, err := x.DBRepo.InsertOnePatient(patient)
	if err != nil {
		controller.RenderFailed(rw, r, err, dir_failed)
		return
	}

	controller.RenderPage(rw, r, id.String(), dir_success)
}
