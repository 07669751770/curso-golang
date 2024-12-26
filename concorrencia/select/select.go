package main

import (
	"fmt"
	"time"

	titulo "github.com/07669751770/curso-golang/titulo"
)

func oMaisRapido(u1, u2, u3, u4 string) string {
	c1 := titulo.TituloStream(u1)
	c2 := titulo.TituloStream(u2)
	c3 := titulo.TituloStream(u3)
	c4 := titulo.TituloStream(u4)

	select {
	case <-c1:
		return u1
	case <-c2:
		return u2
	case <-c3:
		return u3
	case <-c4:
		return u4
	case <-time.After(time.Millisecond * 1000):
		return "Todos perderam!"
	}
}
func main() {

	campeao := oMaisRapido("https://www.google.com", "https://www.github.com", "https://www.stackoverflow.com", "https://www.reddit.com")
	fmt.Println(campeao)

}
