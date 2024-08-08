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

}

func CriaProdutosHandler(w http.ResponseWriter, r *http.Request) {

}
