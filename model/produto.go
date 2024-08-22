package model

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
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

func BuscaPorNome(nomeProduto string) Produto {
	db := db.ConectaBancoDados()

	res := db.QueryRow("SELECT * FROM produtos where nome = $1", nomeProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	if err == sql.ErrNoRows {
		fmt.Printf("Produto nao encontrado %s\n", nome)

	} else {
		panic(err.Error())
	}

	var produto1 Produto
	produto1.ID = id
	produto1.Nome = nome
	produto1.Descricao = descricao
	produto1.Preco = preco
	produto1.Imagem = imagem
	produto1.QuantidadeEmEstoque = quantidade

	defer db.Close()
	return produto1

}

func CriaProduto(prod Produto) error {

	if produtoCadastrado(prod.Nome) {

		fmt.Printf("Produto já cadastrado: %s\n", nome)
		return fmt.Errorf("Produto ja cadastrado")
	}

	db := db.ConectaBancoDados()
	id := uuid.NewString()
	nome := prod.Nome
	preco := prod.Preco
	descricao := prod.Descricao
	imagem := prod.Imagem
	quantidade := prod.QuantidadeEmEstoque

	strInsert := "INSERT INTO produtos VALUES($1, $2, $3, $4, $5, $6)"

	result, err := db.Exec(strInsert, id, nome, strconv.FormatFloat(preco, 'f', 1, 64), descricao, imagem, strconv.Itoa(quantidade))

	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Produto %s criado com sucesso (%d row affected)\n", id, rowsAffected)

	defer db.Close()

	return nil
}

func produtoCadastrado(nomeProduto string) bool {

	prod := BuscaPorNome(nomeProduto)

	return prod.Nome == nomeProduto

}
