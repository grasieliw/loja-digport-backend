package main

import (
	"encoding/json"
	"net/http"

	"github.com/grasieliw/loja-digport-backend/blob/main/main.go/model"
)

func StarServer() {

	http.HandleFunc("/produtos", produtosHandler)
	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		getProdutos(w, r)
	} else if r.Method == "POST" {
		addProdutos(w, r)
	}

}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("nome")

	if queryParam != "" {
		prodBusca := buscaPorNome(queryParam)
		json.NewEncoder(w).Encode(prodBusca)
	} else {
		produtos := catalogo()
		json.NewEncoder(w).Encode(produtos)
	}
}

func addProdutos(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	registrarProdutos(produto)

	w.WriteHeader(http.StatusCreated)

	jsonData, err := json.Marshal(produto)

	if err != nil {
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
	}

	w.Write(jsonData)
}
