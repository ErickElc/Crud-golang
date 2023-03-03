package routes

import (
	"net/http"

	"WebAppGolang/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.IndexPage)
	http.HandleFunc("/New", controllers.NewPage)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
}
