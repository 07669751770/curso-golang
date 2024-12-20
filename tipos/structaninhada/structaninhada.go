package main

import "fmt"

type Pedido struct {
	Cliente string
	Numero  uint64
	Itens   []Item
}

type Item struct {
	Quantidade uint
	Descricao  string
	Valor      float32
}

func main() {

	pedido := Pedido{
		Cliente: "Bruno Pueyo",
		Numero:  1,
		Itens: []Item{
			{
				Quantidade: 2,
				Descricao:  "Camiseta",
				Valor:      29.90,
			},
			{
				Quantidade: 1,
				Descricao:  "Calça Jeans",
				Valor:      99.90,
			},
		},
	}
	pedido.AdicionarItem(Item{3, "Sapato", 129.90})

	total := pedido.ValorTotal()
	fmt.Printf("Valor total do pedido: R$ %.2f\n", total)

}

func (p *Pedido) AdicionarItem(item Item) {
	p.Itens = append(p.Itens, item)
}

func (p Pedido) ValorTotal() float32 {
	total := float32(0)
	for _, item := range p.Itens {
		total += item.Valor * float32(item.Quantidade)
	}
	return total
}
