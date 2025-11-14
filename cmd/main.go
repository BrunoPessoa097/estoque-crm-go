package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/brunopessoa097/estudos-geral/historico"
	"github.com/brunopessoa097/estudos-geral/produto"
	"github.com/brunopessoa097/estudos-geral/utils"
)

func adicionar() {
	reader := bufio.NewReader(os.Stdin)
	var novoProdut = utils.Produto{}

	utils.Clear()

	fmt.Println("--- Adicionando Produto ---")
	// nome do produto
	fmt.Println("Nome do produto: ")
	nome, _ := reader.ReadString('\n')
	novoProdut.Nome = strings.TrimSpace(nome)
	// fim

	// preço
	fmt.Println("Preço do produto: ")
	preco, _ := reader.ReadString('\n')
	preco = strings.TrimSpace(preco)
	precoStr, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		fmt.Println("Preço inválido, usando 0.0")
		precoStr = 0.0
	}
	novoProdut.Preco = precoStr
	// fim preco

	// estoque
	fmt.Println("Estoque: ")
	estoque, _ := reader.ReadString('\n')
	estoque = strings.TrimSpace(estoque)
	estoqueStr, err := strconv.Atoi(estoque)
	if err != nil {
		fmt.Println("Estoque inválido, usando 0")
		estoqueStr = 0
	}
	novoProdut.Estoque = int(estoqueStr)
	// fim estoque

	// adicionando produto
	fmt.Println(produto.Add(novoProdut.Nome, novoProdut.Preco, novoProdut.Estoque))

	menu2(adicionar)
}

func listar() {
	utils.Clear()
	fmt.Println("--- Lista de Produto(s) ---")
	produto.Listar()
	fmt.Println("------")

	menu2(listar)
}

func vender() {
	utils.Clear()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Vender produtos ---")
	fmt.Print("Produto: ")
	prod, _ := reader.ReadString('\n')
	prod = strings.TrimSpace(prod)

	fmt.Println("Quantidade: ")
	estoqueStr, _ := reader.ReadString('\n')
	estoqueStr = strings.TrimSpace(estoqueStr)
	quant, err := strconv.Atoi(estoqueStr)

	if err != nil {
		fmt.Println("Estoque inválido")
		quant = 0
	}

	str, _ := produto.Vender(prod, quant)
	fmt.Println(str)

	menu2(vender)
}

func hist() {
	utils.Clear()
	fmt.Println("--- Histórico ---")
	historico.Listar()
	fmt.Println("------")

	menu2(hist)
}

func menu2(f func()) {
	var escolha int
	fmt.Println("Voltar para o menu?")
	fmt.Println("[1]- Sim")
	fmt.Println("[2]- Não")
	fmt.Scanln(&escolha)

	if escolha == 1 {
		menu()
	} else {
		f()
	}
}

func atualiz() {
	utils.Clear()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- Adicionando estoque ao produto ---")
	fmt.Printf("Produto: ")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)

	fmt.Println("Quantidade: ")
	estoqueStr, _ := reader.ReadString('\n')
	estoqueStr = strings.TrimSpace(estoqueStr)
	quant, err := strconv.Atoi(estoqueStr)

	if err != nil {
		fmt.Println("Estoque inválido")
		quant = 0
	}

	str, _ := produto.AtualizarProduto(nome, quant)
	fmt.Println(str)

	menu2(atualiz)
}

func menu() {
	utils.Clear()
	var escolha int

	fmt.Println("--- Menu ---")
	fmt.Println("[1]- Adicionar produto")
	fmt.Println("[2]- Listar produtos")
	fmt.Println("[3]- Vender produto")
	fmt.Println("[4]- Adicionar Estoque ao Produto")
	fmt.Println("[5]- Histórico")
	fmt.Println("[0]- Sair")
	fmt.Println("Qual sua escolha?")
	fmt.Scanln(&escolha)

	switch escolha {
	case 1:
		adicionar()
	case 2:
		listar()
	case 3:
		vender()
	case 4:
		atualiz()
	case 5:
		hist()
	default:
		fmt.Println("Obrigado por usar !")
	}
}

func main() {
	utils.Clear()
	menu()
}
