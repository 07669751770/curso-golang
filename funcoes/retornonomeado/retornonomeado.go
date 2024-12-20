package main

import "fmt"

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}
