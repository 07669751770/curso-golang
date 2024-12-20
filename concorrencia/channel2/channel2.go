package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // Enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base

	close(c) // Fechando o canal
}

func main() {

	// Channels são a forma de comunicação entre goroutines

	c := make(chan int)

	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // Recebendo dados do canal (bloqueante)
	fmt.Println("B")
	fmt.Println(a, b)
	fmt.Println("C")
	fmt.Println(<-c) // Recebendo o último dado do canal

}
