package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Operação bloqueante
	fmt.Println("Só que depois que foi lido")
}
func main() {

	c := make(chan int) // Canal sem buffer

	go rotina(c)

	fmt.Println(<-c)
	time.Sleep(2 * time.Second)
	fmt.Println("Foi lido")
	// fmt.Println(<-c) // Deadlock
	// fmt.Println("FIM")

	// close(c)
}
