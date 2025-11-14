package historico

import (
	"fmt"

	"github.com/brunopessoa097/estudos-geral/utils"
)

func Add(tipo string, p *utils.Produto) {
	utils.Hist = append(utils.Hist, utils.Historico{
		Tipo:    tipo,
		Produto: p,
	})
}

func Listar() {
	if len(utils.Hist) == 0 {
		fmt.Println("Não há histórico")
	}

	for _, res := range utils.Hist {
		//fmt.Println(res)
		fmt.Printf("Tipo: %s | Nome: %s | Preço: %.2f | Estoque: %d \n", res.Tipo, res.Produto.Nome, res.Produto.Preco, res.Produto.Estoque)
	}
}
