package model

type Carrinho struct {
	ID            string          `json:"id"`
	UserID        string          `json:"userid"`
	InfosProdutos []InfosProdutos `json:"infosprodutos"`
	ValorTotal    float32         `json:"valortotal"`
}

type InfosProdutos struct {
	ProdutoID           string
	QuantidadeDeProduto int
}
