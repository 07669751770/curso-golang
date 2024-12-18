package main

func imprimeResultado(nota float64) {
	if nota >= 7 {
		println("Aprovado com nota", nota)
	} else {
		println("Reprovado com nota", nota)
	}
}

func main() {
	imprimeResultado(7.3)
	imprimeResultado(5.1)
}
