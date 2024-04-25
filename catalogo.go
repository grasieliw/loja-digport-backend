package main

import (
	"errors"

	"github.com/grasieliw/loja-digport-backend/blob/main/main.go/model"
)

var catalogoVazio []model.Produto = []model.Produto{}

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
		{
			Nome:       "Agua",
			Descricao:  "Agua sem gas",
			Categoria:  "Bebidas",
			Id:         "002b",
			Valor:      5,
			Quantidade: 4,
			Imagem:     "imagem.jpg",
		},
		{
			Nome:       "Corona",
			Descricao:  "Cerveja",
			Categoria:  "Bebidas",
			Id:         "003b",
			Valor:      15,
			Quantidade: 10,
			Imagem:     "imagem.jpg",
		},
	}
	return item
}

func buscaPorNome(nome string) []model.Produto {

	produtos := catalogo()
	var produtoEscolhidoNome []model.Produto

	for i := range produtos {
		produtoAtual := produtos[i]
		if produtoAtual.Nome == nome {
			produtoEscolhidoNome = append(produtoEscolhidoNome, produtoAtual)
		}
	}
	return produtoEscolhidoNome

}

func registrarProdutos(novoProduto model.Produto) error {

	for _, produtoDaLista := range catalogoVazio {

		if novoProduto.Id == produtoDaLista.Id {

			return errors.New("produto ja cadastrado")
		}
	}
	catalogoVazio = append(catalogoVazio, novoProduto)
	return nil
}
