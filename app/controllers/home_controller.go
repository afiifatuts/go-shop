package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to Go Shop Home Page")
	//render ke html
	render := render.New(render.Options{Layout: "Layout"})

	//nama file home
	// lempar data dengan variable
	//secara default render akan mencari folder template dan sejajar dengan app
	//extentionnya tmpl
	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "Home Title",
		"body":  "Home Description",
	})
}
