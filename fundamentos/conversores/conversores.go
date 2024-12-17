package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(123)) // converte para o valor ASCII

	// Inteiro para string
	fmt.Println("Teste " + fmt.Sprint(123))
	fmt.Println("Teste", 123)
	fmt.Println("Teste " + fmt.Sprintf("%d", 123))
	fmt.Println("Teste " + strconv.Itoa(123))

	// String para inteiro
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // 1, t, T, TRUE, true, True
	if b {
		fmt.Println("Verdadeiro")
	}

}
