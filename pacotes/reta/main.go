package main

import (
	"fmt"
	"math"
)

func main() {
	ponto1 := Ponto{2, 3}
	ponto2 := Ponto{X: 3, Y: 4}

	fmt.Println(catetos(ponto1, ponto2))
	fmt.Println(Distancia(ponto1, ponto2))
}

// Catetos é responsável por calcular os catetos do triângulo formado por dois pontos
func catetos(p1, p2 Ponto) (float64, float64) {
	catetoX := math.Abs(p2.X - p1.X)
	catetoY := math.Abs(p2.Y - p1.Y)
	return catetoX, catetoY
}

// Distancia é responsável por calcular a distância linear entre dois pontos no plano cartesiano
func Distancia(p1, p2 Ponto) float64 {
	catX, catY := catetos(p1, p2)
	return math.Sqrt(math.Pow(catX, 2) + math.Pow(catY, 2))
}
