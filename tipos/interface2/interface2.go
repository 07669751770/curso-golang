package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Carro struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (c *Carro) ligarTurbo() {
	c.turboLigado = true
}

func main() {

	ferrari := Carro{"F40", false, 0}
	ferrari.ligarTurbo()

	civic := Carro{"Civic", false, 0}

	var sanderoRS Esportivo = &Carro{"Sandero RS", false, 0}
	fmt.Println(ferrari)
	fmt.Println(civic)
	sanderoRS.ligarTurbo()
	fmt.Println(sanderoRS)

}
