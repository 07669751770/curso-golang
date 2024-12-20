package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
		"Rafael Silva":   1350.00,
		"Lucas Silva":    1200.00,
	}

	imprimirMap(funcESalarios)
	funcESalarios["Bruno Pueyo"] = 13050.00
	imprimirMap(funcESalarios)
}

func imprimirMap(m map[string]float64) {
	fmt.Println("---------------")
	for nome, salario := range m {
		fmt.Printf("%s => %.2f\n", nome, salario)
	}
}
