package main

import "fmt"

func main() {
	var b byte = 3
	println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	println(i)

	x, y := 1, 2
	println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
