package main

import (
	"fmt"

	titulo "github.com/07669751770/curso-golang/titulo"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	// Primeiro array com URLs
	urlsGroup1 := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
	}

	// Segundo array com URLs
	urlsGroup2 := []string{
		"https://www.wikipedia.org",
		"https://www.medium.com",
		"https://www.amazon.com",
		"https://www.microsoft.com",
	}

	c := juntar(titulo.Titulo(urlsGroup1), titulo.Titulo(urlsGroup2))

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)

}
