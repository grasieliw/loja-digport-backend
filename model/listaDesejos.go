package model

type listaDesejos struct {
	ID        string   `json:"id"`
	ProdutoID []string `json:"produtoid"`
	UserID    string   `json:"userid"`
}
