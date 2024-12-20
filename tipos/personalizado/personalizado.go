package main

import "fmt"

type Nota float64

func (n Nota) isBetween(notaInicio float64, notaFim float64) bool {
	return float64(n) >= notaInicio && float64(n) <= notaFim
}

func notaParaConceito(n Nota) string {
	if n.isBetween(9.0, 10.0) {
		return "A"
	} else if n.isBetween(7.0, 8.9) {
		return "B"
	} else if n.isBetween(5.0, 6.9) {
		return "C"
	} else if n.isBetween(3.0, 4.9) {
		return "D"
	}
	return "E"
}

func main() {

	var nota Nota
	nota = 9.8
	fmt.Println(notaParaConceito(nota))

	nota = 6.9
	fmt.Println(notaParaConceito(nota))

	nota = 2.1
	fmt.Println(notaParaConceito(nota))

}
