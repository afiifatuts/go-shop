package controllers

import (
	"net/http"
	"strconv"

	"github.com/afiifatuts/go-shop/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to Go Shop Home Page")
	//render ke html
	render := render.New(render.Options{Layout: "Layout"})

	//mengmbil query url misal localhost:9000/products?page=
	q := r.URL.Query()
	// mengambil value/key dari pagenya ?page=1
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage := 9

	//memangil model
	productModel := models.Product{}
	//memanggil method find
	products, totalRows, err := productModel.GetProducts(server.DB, perPage, page)
	if err != nil {
		return
	}

	//nama file home
	// lempar data dengan variable
	//secara default render akan mencari folder template dan sejajar dengan app
	//extentionnya tmpl
	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products": products,
	})
}
