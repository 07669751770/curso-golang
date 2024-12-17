package main

import "fmt"

func main() {

	var a int
	var b float64
	var c bool
	var d string
	var e *int

	f := 2
	e = &f

	fmt.Printf("a: %v b:%v c:%v d:%v *e:%v &e:%v", a, b, c, d, *e, &e)

}
