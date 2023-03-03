package models

import (
	"WebAppGolang/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Produto {
	db := db.ConnectDataBase()
	selectProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}
	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func BuscarItemId(Id int) Produto {
	db := db.ConnectDataBase()

	produtoDb, err := db.Query("select * from produtos where id=$1", Id)

	if err != nil {
		panic(err.Error())
	}

	produtoSalvo := Produto{}

	for produtoDb.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDb.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoSalvo.Id = id
		produtoSalvo.Nome = nome
		produtoSalvo.Descricao = descricao
		produtoSalvo.Preco = preco
		produtoSalvo.Quantidade = quantidade

	}
	defer db.Close()
	return produtoSalvo
}

func CriarNovoProduto(name, descricao string, preco float64, quantidade int) {
	db := db.ConnectDataBase()

	insertDataBase, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}
	insertDataBase.Exec(name, descricao, preco, quantidade)
	defer db.Close()
}

func EditarProduto(name, descricao string, preco float64, quantidade, id int) {
	db := db.ConnectDataBase()

	updateDataBase, err := db.Prepare("update produtos set nome=$2, descricao=$3, preco=$4, quantidade=$5 where id=$1")

	if err != nil {
		panic(err.Error())
	}
	updateDataBase.Exec(id, name, descricao, preco, quantidade)
	defer db.Close()
}

func DeletarProduto(IdProduto int) {
	db := db.ConnectDataBase()

	deleteProduct, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(IdProduto)

	defer db.Close()
}
