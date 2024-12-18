package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a    // b é um ponteiro para a
	fmt.Println(a, *b) // *b é o valor de a
	a = 27             // altera o valor de a
	fmt.Println(a, *b) // *b é o valor de a
	*b = 14            // altera o valor de a
	fmt.Println(a, *b) // *b é o valor de a
	fmt.Println(&a, b) // &a é o endereço de a, b é o endereço de a
}
