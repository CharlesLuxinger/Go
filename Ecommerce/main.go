package main

import (
	"html/template"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"Go/Ecommerce/produtos"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := dbConnect()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func dbConnect() *sql.DB {
	var db, err = sql.Open("mysql", "root:12345678@/ecommerce?tls=skip-verify&autocommit=true&charset=utf8mb4")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()

	allProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := produtos.Produto{}
	produtos := []produtos.Produto{}

	for allProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = allProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	templates.ExecuteTemplate(w, "Index", allProdutos)

	defer db.Close()
}
