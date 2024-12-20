package main

import (
	"fmt"
	"time"
)

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func primeNumbers(n int, c chan int) {
	inicio := 2
	for i := 0; i <= n; i++ {
		for primo := inicio; ; primo++ {
			if isPrime(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}

	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primeNumbers(cap(c), c)
	for prime := range c {
		fmt.Printf("%d ", prime)
	}
}
