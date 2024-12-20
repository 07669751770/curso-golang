package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[987654321] = "Pedro"
	aprovados[456789123] = "Ana"
	fmt.Println("Imprimindo map")
	PrintMap(aprovados)
	fmt.Println("Imprimindo um registro do map")
	fmt.Println(aprovados[123456789])
	fmt.Println("Deletando um registro do map")
	delete(aprovados, 123456789)
	fmt.Println("Imprimindo map")
	PrintMap(aprovados)
}

func PrintMap(m map[int]string) {

	for cpf, nome := range m {
		fmt.Printf("%s CPF[%d]\n", nome, cpf)
	}
}
