package produto

import (
	"fmt"

	"github.com/brunopessoa097/estudos-geral/historico"
	"github.com/brunopessoa097/estudos-geral/model"
)

var produtos = make(map[string]model.Produto)

func Add(nome string, preco float64, estoque int) string {
	produtos[nome] = model.Produto{
		Nome:    nome,
		Preco:   preco,
		Estoque: estoque,
	}
	historico.Add("Adicionado", &model.Produto{
		Nome:    nome,
		Preco:   preco,
		Estoque: estoque,
	})
	return "Produto adicionado"
}

func Listar() {
	if len(produtos) == 0 {
		fmt.Println("Sem Produtos para mostrar")
	}

	for _, prod := range produtos {
		fmt.Printf("Nome: %s | Preço: %.2f | Estoque: %d \n", prod.Nome, prod.Preco, prod.Estoque)
	}
}

func Vender(nome string, quant int) (string, bool) {
	prod, ok := produtos[nome]

	if !ok {
		return "Produto não existe", false

	}

	if quant > int(prod.Estoque) {
		return "Quantidade maior que exitem em estoque", false

	}

	prod.Estoque -= int(quant)
	produtos[nome] = prod

	historico.Add("Venda", &model.Produto{
		Nome:    nome,
		Preco:   prod.Preco,
		Estoque: quant,
	})
	return fmt.Sprintf("Venda efetuada do %s com sucesso", nome), true
}

func AtualizarProduto(nome string, adicionado int) (string, bool) {
	prod, ok := produtos[nome]

	if !ok {
		return "Produto não existe", false
	}

	prod.Estoque += adicionado
	produtos[nome] = prod

	historico.Add("Atualizando estoque", &model.Produto{
		Nome:    nome,
		Preco:   prod.Preco,
		Estoque: adicionado,
	})

	return fmt.Sprintf("Produto %s | Adicionado ao estoque: %d", prod.Nome, adicionado), true
}
