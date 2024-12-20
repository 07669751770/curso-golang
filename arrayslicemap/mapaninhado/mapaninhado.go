package main

import "fmt"

func main() {
	funcESalarios := map[string]map[string]float64{
		"José João": {
			"Salário": 11325.45,
			"CPF":     123456789,
		},
		"Gabriela Silva": {
			"Salário": 15456.78,
			"CPF":     987654321,
		},
		"Pedro Junior": {
			"Salário": 1200.00,
			"CPF":     456789123,
		},
		"Rafael Silva": {
			"Salário": 1350.00,
			"CPF":     456789124,
		},
		"Lucas Silva": {
			"Salário": 1200.00,
			"CPF":     456789125,
		},
	}

	imprimirMap(funcESalarios)
	funcESalarios["Bruno Pueyo"] = map[string]float64{
		"Salário": 13050.00,
	}
	imprimirMap(funcESalarios)
}

func imprimirMap(m map[string]map[string]float64) {
	fmt.Println("---------------")
	for nome, dados := range m {
		fmt.Printf("%s\n", nome)
		for chave, valor := range dados {
			fmt.Printf("---%s => %.2f\n", chave, valor)
		}
	}
}
