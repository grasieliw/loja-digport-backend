package model

import (
	"database/sql"

	"github.com/grasieliw/loja-digport-backend/db"
)

type Produto struct {
	ID                  string  `json:"id"`
	Nome                string  `json:"nome"`
	Preco               float64 `json:"preco"`
	Descricao           string  `json:"descricao"`
	Imagem              string  `json:"imagem"`
	QuantidadeEmEstoque int     `json:"quantidade"`
}

var id, nome string
var preco float64
var descricao, imagem string
var quantidade int

func BuscaProdutos() []Produto {

	db := db.ConectaBancoDados()

	resultado, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for resultado.Next() {

		err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Imagem = imagem
		p.QuantidadeEmEstoque = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func BuscaPorNome(){
	db := db.ConectaBancoDados()

	res := db.QueryRow("SELECT * FROM produtos where nome = $1", nomeProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	 if err = sql.ErrNoRows{
		fmt.Printf("Produto noa encontrado %s\n", nome)
	
	 } else {
		panic(err.Error())
	 }

	 var produto1 Produto
	 p.ID = id
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Imagem = imagem
	p.QuantidadeEmEstoque = quantidade

	defer db.Close()
	return produto1

}