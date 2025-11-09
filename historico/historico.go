package historico

import (
	"fmt"

	"github.com/brunopessoa097/estudos-geral/model"
)

var hist = make([]model.Historico, 0)

func Add(tipo string, p *model.Produto) {
	hist = append(hist, model.Historico{
		Tipo:    tipo,
		Produto: p,
	})
}

func Listar() {
	if len(hist) == 0 {
		fmt.Println("Não há histórico")
	}

	for _, res := range hist {
		//fmt.Println(res)
		fmt.Printf("Tipo: %s | Nome: %s | Preço: %.2f | Estoque: %d \n", res.Tipo, res.Produto.Nome, res.Produto.Preco, res.Produto.Estoque)
	}
}
