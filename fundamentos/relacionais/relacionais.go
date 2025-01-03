package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("d1 > d2", d1.After(d2))
	fmt.Println("d1 == d2", d1 == d2)
	fmt.Println("d1 == d2", d1.Equal(d2))
	fmt.Println("d1 < d2", d1.Before(d2))

	type Pessoa struct {
		nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("p1 == p2", p1 == p2)

	p2 = Pessoa{"Joã"}
	fmt.Println("p1 == p2", p1 == p2)
}
