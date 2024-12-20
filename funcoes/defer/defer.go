package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("obterValorAprovado - Fim!")
	if numero > 5000 {
		fmt.Println("obterValorAprovado - Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println("main", obterValorAprovado(6000))
}
