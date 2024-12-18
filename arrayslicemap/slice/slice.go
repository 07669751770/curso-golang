package main

import (
	"fmt"
	"reflect"
)

func main() {

	notas := []float64{7.8, 4.3, 9.1}   // slice
	notas2 := [3]float64{7.8, 4.3, 9.1} // array

	fmt.Printf("tipo(notas): %v\ntipo(notas2): %v\n", reflect.TypeOf(notas), reflect.TypeOf(notas2))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // slice
	fmt.Printf("tipo(a2): %v\ntipo(s2): %v\n", reflect.TypeOf(a2), reflect.TypeOf(s2))

	s2[1] = 10 // altera o valor do array

	s3 := a2[:2]            // slice
	fmt.Println(a2, s2, s3) // [1 2 10 4 5] [2 10] [1 2]
	s3[1] = 20              // altera o valor do array
	fmt.Println(a2, s2, s3) // [1 20 10 4 5] [20 10] [1 20]
}
