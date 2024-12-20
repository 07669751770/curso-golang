package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}
type Luxuoso interface {
	fazerBaliza()
}

type EsportivoLuxuoso interface {
	Esportivo
	Luxuoso
	configuraSuspensao()
}

// type bmwSerie7 struct {
// 	modelo          string
// 	turboLigado     bool
// 	velocidadeAtual int
// 	suspensao       string
// }

type bmwSerie7 struct{}

func (b *bmwSerie7) ligarTurbo() {
	fmt.Println("Turbo ligado")
}
func (b *bmwSerie7) fazerBaliza() {
	fmt.Println("Baliza feita")
}
func (b *bmwSerie7) configuraSuspensao() {
	fmt.Println("Suspensão configurada")
}

func ligarTurbo(veiculo Esportivo) {
	veiculo.ligarTurbo()
}

func fazerBaliza(veiculo Luxuoso) {
	veiculo.fazerBaliza()
}

func configurarSuspensao(veiculo EsportivoLuxuoso) {
	veiculo.configuraSuspensao()
}

func main() {

	serie7 := &bmwSerie7{}
	ligarTurbo(serie7)
	fazerBaliza(serie7)
	configurarSuspensao(serie7)
}
