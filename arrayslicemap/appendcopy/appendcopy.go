package main

import (
	"fmt"
	"unsafe"
)

func main() {

	array1 := [3]int{1, 2, 3}
	var slice1 = array1[:]

	fmt.Println(array1, slice1)
	fmt.Println(unsafe.Pointer(&array1[0]), unsafe.Pointer(&slice1[0]))

	slice1 = append(slice1, 4)

	fmt.Println(array1, slice1)
	fmt.Println(unsafe.Pointer(&array1[0]), unsafe.Pointer(&slice1[0]))

	slice1[0] = 10 // altera o array
	fmt.Println(array1, slice1)
	fmt.Println(unsafe.Pointer(&array1[0]), unsafe.Pointer(&slice1[0]))

	slice2 := make([]int, len(slice1), cap(slice1))
	copy(slice2, slice1)
	fmt.Println(array1, slice1, slice2)
	fmt.Println(unsafe.Pointer(&array1[0]), unsafe.Pointer(&slice1[0]), unsafe.Pointer(&slice2[0]))
}
