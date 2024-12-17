package main

import (
	"fmt"
	"math"
)

func main() {
	//const PI float64 = 3.1415

	const (
		PI    = 3.1415
		NUM_2 = 2
	)
	var raio = 3.2 // tipo (float64) inferido pelo compilador
	area := PI * math.Pow(raio, NUM_2)

	fmt.Printf("A área da circunferência é %.2f\n", area)

	var h, g, i = "teste", false, 3.7

	fmt.Println(h, g, i)

	texto, e, f := "numero", 2, 4
	fmt.Println(texto, e, f)
}
