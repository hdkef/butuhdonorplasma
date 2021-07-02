package main

import (
	"butuhdonorplasma/dbdriver"
	"butuhdonorplasma/handler"
	"butuhdonorplasma/public"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/NYTimes/gziphandler"
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

var DBNAME string
var PORT string

func init() {
	godotenv.Load()
	PORT = os.Getenv("PORT")
	DBNAME = os.Getenv("DBNAME")
}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	dbclient, err := dbdriver.DBConn(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	dbrepo := dbdriver.DBRepo{
		DB: dbclient.Database(DBNAME),
	}
	defer dbclient.Disconnect(ctx)

	publicpages := public.GetPublicPages(&dbrepo)

	r := mux.NewRouter()
	r.HandleFunc("/", publicpages.IndexPage())
	r.HandleFunc("/find", publicpages.FindPage())
	r.HandleFunc("/add", publicpages.AddPage())
	r.HandleFunc("/delete", publicpages.DeletePage())
	r.HandleFunc("/result", publicpages.ResultPage())
	r.HandleFunc("/getcity", handler.GetCityHandler())

	public := publicHandler{staticPath: os.Getenv("STATICPATH")}
	r.PathPrefix("/").HandlerFunc(public.ServeHTTP)

	addr := fmt.Sprintf(":%s", PORT)

	fmt.Println("listening to port : ", PORT)

	err = http.ListenAndServe(addr, gziphandler.GzipHandler(r))
	if err != nil {
		panic(err.Error())
	}
}
