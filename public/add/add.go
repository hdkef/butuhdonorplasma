package add

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"butuhdonorplasma/konstant"
	"butuhdonorplasma/models"
	"embed"
	"errors"
	"html/template"
	"net/http"
	"strconv"
)

//go:embed *.html
var addtmpl embed.FS

var addtemplates *template.Template = template.Must(template.ParseFS(addtmpl, "add.html"))
var addsuccesstemplates *template.Template = template.Must(template.ParseFS(addtmpl, "add-success.html"))
var addfailtemplates *template.Template = template.Must(template.ParseFS(addtmpl, "fail.html"))

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

		controller.RenderPage(rw, r, data, addtemplates)
	}
}

func handleAddPost(rw http.ResponseWriter, r *http.Request, x *AddHandler) {

	err := r.ParseForm()
	if err != nil {
		controller.RenderFailed(rw, r, err, addfailtemplates)
		return
	}

	captchaID, err := strconv.Atoi(r.FormValue("captchaid"))
	if err != nil {
		controller.RenderFailed(rw, r, err, addfailtemplates)
		return
	}
	captcha := r.FormValue("captcha")

	err = controller.CheckCaptcha(int64(captchaID), captcha)
	if err != nil {
		controller.RenderFailed(rw, r, err, addfailtemplates)
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
		Cpname1:      r.FormValue(konstant.Cpname1),
		Cptel1:       r.FormValue(konstant.Cptel1),
		Cprelation1:  r.FormValue(konstant.Cprelation1),
	}

	if patient.Name == "" || patient.Age == "" || patient.Desc == "" || patient.HospitalName == "" || patient.ProvinceID == "" || patient.CityID == "" || patient.Goldar == "" || patient.Rhesus == "" {
		controller.RenderFailed(rw, r, errors.New("something is missing"), addfailtemplates)
		return
	}

	//ADD TO DATABASE
	id, err := x.DBRepo.InsertOnePatient(patient)
	if err != nil {
		controller.RenderFailed(rw, r, err, addfailtemplates)
		return
	}

	controller.RenderPage(rw, r, id.String(), addsuccesstemplates)
}
