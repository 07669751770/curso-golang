package main

import (
	"fmt"
	"unsafe"
)

func main() {

	s := []int{1, 2, 3, 4, 5, 6}

	s2 := append(s, 7, 8, 9, 0)

	fmt.Println(s, s2)

	s2[0] = 1000 // altera o valor do array
	fmt.Println(s, s2)

	s3 := s2[:] // slice
	fmt.Println(s, s2, s3)

	s2[1] = 2000 // altera o valor do array
	fmt.Println(s, s2, s3)

	fmt.Println("-------------------------------------------------------")

	array := [5]int{10, 20, 30, 40, 50}
	slice := array[:]

	ptrSlice := unsafe.Pointer(&slice[0])

	// Endereço do array no índice 1
	ptrArray := unsafe.Pointer(&array[0])

	// Exibindo o ponteiro e outros dados
	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Ponteiro.Array: %x\n", ptrArray) // Endereço base do array subjacente
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Ponteiro.slice: %x\n", ptrSlice) // Endereço base do array subjacente
	fmt.Printf("Length: %d\n", len(slice))
	fmt.Printf("Capacity: %d\n", cap(slice))

	copy(s3, slice)

	prtS3 := unsafe.Pointer(&s3[0])

	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("Ponteiro.s3: %x\n", prtS3) // Endereço base do array subjacente
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Ponteiro.slice: %x\n", ptrSlice) // Endereço base do array subjacente

}
