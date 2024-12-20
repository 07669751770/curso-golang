package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Porque você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de vc!", 1)
	go fale("Maria", "Porque você não fala comigo?", 10)
	go fale("João", "Só posso falar depois de vc!", 500)

}
