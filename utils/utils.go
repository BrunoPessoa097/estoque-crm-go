package utils

import (
	"os"
	"os/exec"
	"runtime"
)

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
	Hist     = make([]Historico, 0)
)

func Clear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windowns" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
