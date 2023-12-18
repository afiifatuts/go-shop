package app

import (
	"github.com/afiifatuts/go-shop/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) initializeRouter() {

	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
