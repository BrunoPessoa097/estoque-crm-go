package produto

import (
	"fmt"

	"github.com/brunopessoa097/estudos-geral/historico"
	"github.com/brunopessoa097/estudos-geral/utils"
)

func Add(nome string, preco float64, estoque int) string {
	utils.Produtos[nome] = utils.Produto{
		Nome:    nome,
		Preco:   preco,
		Estoque: estoque,
	}
	historico.Add("Adicionado", &utils.Produto{
		Nome:    nome,
		Preco:   preco,
		Estoque: estoque,
	})
	return "Produto adicionado"
}

func Listar() {
	if len(utils.Produtos) == 0 {
		fmt.Println("Sem Produtos para mostrar")
	}

	for _, prod := range utils.Produtos {
		fmt.Printf("Nome: %s | Preço: %.2f | Estoque: %d \n", prod.Nome, prod.Preco, prod.Estoque)
	}
}

func Vender(nome string, quant int) (string, bool) {
	prod, ok := utils.Produtos[nome]

	if !ok {
		return "Produto não existe", false

	}

	if quant > int(prod.Estoque) {
		return "Quantidade maior que exitem em estoque", false

	}

	prod.Estoque -= int(quant)
	utils.Produtos[nome] = prod

	historico.Add("Venda", &utils.Produto{
		Nome:    nome,
		Preco:   prod.Preco,
		Estoque: quant,
	})
	return fmt.Sprintf("Venda efetuada do %s com sucesso", nome), true
}

func AtualizarProduto(nome string, adicionado int) (string, bool) {
	prod, ok := utils.Produtos[nome]

	if !ok {
		return "Produto não existe", false
	}

	prod.Estoque += adicionado
	utils.Produtos[nome] = prod

	historico.Add("Atualizando estoque", &utils.Produto{
		Nome:    nome,
		Preco:   prod.Preco,
		Estoque: adicionado,
	})

	return fmt.Sprintf("Produto %s | Adicionado ao estoque: %d", prod.Nome, adicionado), true
}
