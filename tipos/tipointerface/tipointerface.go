package main

import "fmt"

type curso struct {
	nome string
}

type pessoa struct {
	nome string
}

func main() {
	var coisa interface{}

	coisa = pessoa{"Roberto"}

	fmt.Println(coisa)

	coisa = true

	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = pessoa{"Bruno"}

	fmt.Println(coisa2)

	coisa2 = curso{"Golang"}

	fmt.Println(coisa2)

}
