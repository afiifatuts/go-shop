package app

import "github.com/afiifatuts/go-shop/app/controllers"

func (server *Server) initializeRouter() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
