package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch { // switch true
	case t.Hour() < 12: // se for menor que 12 horas
		fmt.Println("Bom dia!")
		break
	case t.Hour() < 18: // se for menor que 18 horas
		fmt.Println("Boa tarde!")
		break
	default: // caso contrário
		fmt.Println("Boa noite!")
		break
	}

	fmt.Println("9.8", notaParaConceito(9.8))
	fmt.Println("6.9", notaParaConceito(6.9))
	fmt.Println("2.1", notaParaConceito(2.1))
	fmt.Println("-1", notaParaConceito(-1))

}
