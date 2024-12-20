package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	fmt.Println("Executed")
}

func main() {
	buffer := make(chan int, 5)

	for i := 0; i <= 10; i++ {
		select {
		case buffer <- i:
			fmt.Printf("Inseriu %d\n", i)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Buffer full")
			fmt.Println("Removing value from buffer")
			fmt.Printf("Value removed: %d\n", <-buffer)
		}

	}
	close(buffer)
}
