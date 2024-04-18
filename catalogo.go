package main

import (
	"github.com/grasieliw/loja-digport-backend/blob/main/main.go/model"
)

func catalogo() []model.Produto {

	item := []model.Produto{

		{
			Nome:       "Coca-Cola",
			Descricao:  "Refrigerante",
			Categoria:  "Bebidas",
			Id:         "001b",
			Valor:      10,
			Quantidade: 3,
			Imagem:     "imagem.jpg",
		},
	}
	return item
}
