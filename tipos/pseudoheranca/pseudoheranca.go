package main

import "fmt"

type Carro struct {
	Nome            string
	VelocidadeAtual int
}

type Ferrari struct {
	Carro
	TurboLigado bool
}

func main() {

	f := Ferrari{}
	f.Nome = "F40"
	f.VelocidadeAtual = 0
	f.TurboLigado = true
	fmt.Printf("A Ferrari %s está com o turbo ligado? %v\n", f.Nome, f.TurboLigado)
}
