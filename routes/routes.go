package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grasieliw/loja-digport-backend/controller"
	"github.com/rs/cors"
)

func HandleRequests() {

	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.BuscaProdutosPorNomeHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.CriaProdutosHandler).Methods("POST")

	route.HandleFunc("/produtos/{id}", controller.RemoveProdutoHandler).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(route)

	http.ListenAndServe(":8080", handler)
	//http://localhost:8080
}
