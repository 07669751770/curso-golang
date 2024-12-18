package main

import "fmt"

func main() {

	numero := [...]int{1, 2, 3, 4, 5, 6}

	for i, n := range numero {
		fmt.Printf("numero[%d] = %d\n", i, n)
	}
}
