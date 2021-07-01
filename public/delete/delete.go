package delete

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"embed"
	"errors"
	"html/template"
	"net/http"
)

type DeleteHandler struct {
	DBRepo *dbdriver.DBRepo
}

//go:embed delete.html
var deletetmpl embed.FS

//go:embed delete-success.html
var deletesuccess embed.FS

//go:embed fail.html
var deletefailed embed.FS

var deletetemplates *template.Template = template.Must(template.ParseFS(deletetmpl, "delete.html"))
var deletesuccesstemplates *template.Template = template.Must(template.ParseFS(deletesuccess, "delete-success.html"))
var deletefailtemplates *template.Template = template.Must(template.ParseFS(deletefailed, "fail.html"))

func GetDeleteHandler(dbrepo *dbdriver.DBRepo) *DeleteHandler {
	return &DeleteHandler{
		DBRepo: dbrepo,
	}
}

func (x *DeleteHandler) Delete() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			handleDeletePost(rw, r, x)
			return
		}

		controller.RenderPage(rw, r, nil, deletetemplates)
	}
}

func handleDeletePost(rw http.ResponseWriter, r *http.Request, x *DeleteHandler) {
	err := r.ParseForm()
	if err != nil {
		controller.RenderFailed(rw, r, err, deletefailtemplates)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		controller.RenderFailed(rw, r, errors.New("NO ID inserted"), deletefailtemplates)
		return
	}

	res, err := x.DBRepo.DeletePatientByID(id)
	if err != nil {
		controller.RenderFailed(rw, r, err, deletefailtemplates)
		return
	}

	controller.RenderPage(rw, r, res, deletesuccesstemplates)
}
