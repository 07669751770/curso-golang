package main

func main() {
	x := 1
	y := 2

	// Go only has postfix operators
	x++ // x += 1 or x = x + 1
	println(x)

	y-- // y -= 1 or y = y - 1
	println(y)

	// Go doesn't have prefix operators
	// ++x
	// --y
}
