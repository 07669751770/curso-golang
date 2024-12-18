package main

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A+"
	} else if nota >= 9 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 7 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else if nota >= 0 && nota < 3 {
		return "E"
	} else {
		return "Nota inválida"
	}

}

func main() {
	println(notaParaConceito(9.8))
	println(notaParaConceito(6.9))
	println(notaParaConceito(2.1))
	println(notaParaConceito(-1))
}
