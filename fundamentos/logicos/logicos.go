package main

import "fmt"

func main() {

	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("tv50 = %t, tv32 = %t, sorvete = %t\n", tv50, tv32, sorvete)

	tv50, tv32, sorvete = comprar(true, false)
	fmt.Printf("tv50 = %t, tv32 = %t, sorvete = %t\n", tv50, tv32, sorvete)

	tv50, tv32, sorvete = comprar(false, true)
	fmt.Printf("tv50 = %t, tv32 = %t, sorvete = %t\n", tv50, tv32, sorvete)

	tv50, tv32, sorvete = comprar(false, false)
	fmt.Printf("tv50 = %t, tv32 = %t, sorvete = %t\n", tv50, tv32, sorvete)

}

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // e
	comprarTv32 := trab1 != trab2    // ou exclusivo
	comprarSorvete := trab1 || trab2 // ou

	return comprarTv50, comprarTv32, comprarSorvete
}
