package area

// Pi é uma proporção numérica definida pela relação entre
const Pi = 3.1416

// Circ calcula a área de uma circunferência
func Circ(raio float64) float64 {
	return Pi * raio * raio
}

// Rect calcula a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
