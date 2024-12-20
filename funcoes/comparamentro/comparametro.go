package main

import "fmt"

func multiplicar(p1 int, p2 int) int {
	return p1 * p2
}

func exec(f func(int, int) int, p1, p2 int) int {
	return f(p1, p2)
}
func main() {
	fmt.Println(exec(multiplicar, 2, 2))

}
