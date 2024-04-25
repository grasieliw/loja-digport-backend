package main

import (
	"encoding/json"
	"net/http"
)

func StarServer() {

	http.HandleFunc("/produtos", produtosHandler)
	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("nome")

	if queryParam != "" {
		prodBusca := buscaPorNome(queryParam)
		json.NewEncoder(w).Encode(prodBusca)
	} else {
		produtos := catalogo()
		json.NewEncoder(w).Encode(produtos)
	}

}
