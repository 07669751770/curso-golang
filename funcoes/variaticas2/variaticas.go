package main

import "fmt"

func exibirAprovados(aprovados ...string) {

	for i, nome := range aprovados {
		fmt.Printf("%d) %s\n", i+1, nome)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Ana", "João", "Carlos", "Bruna"}

	exibirAprovados(aprovados...)
}
