package utils

type Produto struct {
	Nome    string
	Preco   float64
	Estoque int
}

type Historico struct {
	Tipo    string
	Produto *Produto
}

var (
	Produtos = make(map[string]Produto)
)
