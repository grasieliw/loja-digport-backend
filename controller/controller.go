package controller

import (
	"encoding/json"
	"net/http"

	"github.com/grasieliw/loja-digport-backend/model"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaProdutos()
	json.NewEncoder(w).Encode(produtos)
}

func BuscaProdutosPorNomeHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	produto := model.BuscaPorNome(nome)
	json.NewEncoder(w).Encode(produto)

}

func CriaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	error := model.CriaProduto(produto)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func RemoveProdutoHandler(w http.ResponseWriter, r *http.Request) {
}
