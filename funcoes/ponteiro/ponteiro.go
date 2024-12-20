package main

import "fmt"

func inc2p(n *int) {
	*n++
}
func inc2cp(n int) {
	n++
}

func main() {
	n := 1
	inc2cp(n) // por valor
	fmt.Println(n)
	inc2p(&n) // por referência
	fmt.Println(n)
}
