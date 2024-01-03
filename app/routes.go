package app

import (
	"net/http"

	"github.com/afiifatuts/go-shop/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) initializeRouter() {

	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

	//definisi route foldernya
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
