package main

import (
	"WebAppGolang/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)

}


