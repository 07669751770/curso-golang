package main

import (
	"fmt"

	titulo "github.com/07669751770/curso-golang/titulo"
)

func main() {
	titulos := []string{"https://www.google.com", "https://www.facebook.com"}
	titulos2 := []string{"https://www.reddit.com", "https://www.twitter.com"}

	t1 := titulo.Titulo(titulos)
	t2 := titulo.Titulo(titulos2)

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)

}
