package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"WebAppGolang/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func IndexPage(w http.ResponseWriter, r *http.Request) {
	produtos := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", produtos)

}

func NewPage(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	const statusCode = 301
	if r.Method == "POST" {
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("erro conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("erro conversão da quantidade:", err)
		}
		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", statusCode)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	const statusCode = 301

	if r.Method == "POST" {
		IdProduto := r.FormValue("id")
		id, err := strconv.Atoi(IdProduto)
		if err != nil {
			panic(err.Error())
		}

		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("erro conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("erro conversão da quantidade:", err)
		}
		models.EditarProduto(nome, descricao, precoConvertido, quantidadeConvertida, id)
	}
	http.Redirect(w, r, "/", statusCode)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	IdProduto := r.URL.Query().Get("id")

	id, err := strconv.Atoi(IdProduto)
	if err != nil {
		panic(err.Error())
	}

	produto := models.BuscarItemId(id)
	temp.ExecuteTemplate(w, "Edit", produto)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	const statusCode = 301
	IdProduto := r.URL.Query().Get("id")

	id, err := strconv.Atoi(IdProduto)
	if err != nil {
		panic(err.Error())
	}

	models.DeletarProduto(id)

	http.Redirect(w, r, "/", statusCode)
}
