package main

import "fmt"

var soma = func(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a int, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
