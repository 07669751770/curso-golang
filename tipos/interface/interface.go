package main

import "fmt"

type Imprimivel interface {
	ToString() string
}

type Pessoa struct {
	Nome string
}

type Produto struct {
	Descricao string
	Preco     float64
}

func (p Pessoa) ToString() string {
	return p.Nome
}
func (p Produto) ToString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.Descricao, p.Preco)
}

func imprimir(x Imprimivel) {
	fmt.Println(x.ToString())
}

func main() {
	bruno := Pessoa{"Bruno"}
	celular := Produto{"Moto Edge 50 Pro", 1899.99}

	imprimir(bruno)
	imprimir(celular)
}
