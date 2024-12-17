package main

import (
	"fmt"
	"os"
)

type file struct{}

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "Opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v", a, b, c, d)

	// fl := file{}

	// fmt.Fprintf(fl, "Minha idade é %d", 30)

	fmt.Fprintf(os.Stdout, "Minha idade é %d", 30)
}

func (file) Write(p []byte) (n int, err error) {

	os.WriteFile("abc.txt", p, 0644)

	return len(p), nil
}
