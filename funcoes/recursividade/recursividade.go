package main

func factorial(n uint) uint { // unsigned integer type (uint) is used to represent only non-negative numbers
	switch {
	case n == 0:
		return 1
	default:
		return n * factorial(n-1)
	}
}

func main() {
	println(factorial(6))
}
