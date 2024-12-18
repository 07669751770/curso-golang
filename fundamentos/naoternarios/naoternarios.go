package main

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
	// return nota >= 6 ? "Aprovado" : "Reprovado" // Go doesn't have ternary operator

}

func main() {
	println(obterResultado(6.2))
	println(obterResultado(5.1))
}
