package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grasieliw/loja-digport-backend/controller"
)

func HandleRequests() {

	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.BuscaProdutosPorNomeHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.CriaProdutosHandler).Methods("POST")

	http.ListenAndServe(":8080", route)
}
