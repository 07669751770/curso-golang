package main

import "fmt"

func soma(numeros ...int) float64 {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return float64(soma) / float64(len(numeros))
}

func main() {
	fmt.Printf("%.2f\n", soma(2, 3))
	fmt.Printf("%.2f\n", soma(2, 3, 4))
}
