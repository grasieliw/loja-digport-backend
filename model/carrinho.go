package model

type Carrinho struct {
	ID            string
	UserID        string
	InfosProdutos []InfosProdutos
	ValorTotal    float32
}

type InfosProdutos struct {
	ProdutoID           string
	QuantidadeDeProduto int
}
