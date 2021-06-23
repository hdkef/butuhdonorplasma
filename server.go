package main

import (
	"butuhdonorplasma/handler"
	"butuhdonorplasma/public"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type publicHandler struct {
	staticPath string
}

//ServeHTTP is to implement interface so that PublicHandler can be use as Handler
func (h publicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, r.URL.Path)

	_, err = os.Stat(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func init() {
	godotenv.Load()
}

func main() {

	publicpages := public.GetPublicPages()

	r := mux.NewRouter()
	r.HandleFunc("/", publicpages.IndexPage())
	r.HandleFunc("/find", publicpages.FindPage())
	r.HandleFunc("/add", publicpages.AddPage())
	r.HandleFunc("/result", publicpages.ResultPage())
	r.HandleFunc("/getcity", handler.GetCityHandler())

	public := publicHandler{staticPath: os.Getenv("STATICPATH")}
	r.PathPrefix("/").HandlerFunc(public.ServeHTTP)

	PORT := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", PORT)

	fmt.Println("listening to port : ", PORT)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		panic(err.Error())
	}
}
