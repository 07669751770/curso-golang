package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("OU =>", a|b)
	fmt.Println("XOR =>", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", c > d) // true
	fmt.Println("Menor =>", c < d) // false

	// Outros operadores

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(float64(a), float64(b)))
	fmt.Println("Exponenciação =>", math.Pow(float64(a), float64(b)))
	fmt.Println("Exponenciação =>", math.Pow(float64(b), float64(a)))

}
