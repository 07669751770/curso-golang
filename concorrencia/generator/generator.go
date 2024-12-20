package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func titulo(urls []string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			re := r.FindStringSubmatch(string(html))
			if len(re) >= 2 {
				c <- r.FindStringSubmatch(string(html))[1]
			} else {
				c <- "Titulo ausente"
			}

		}(url)
	}
	return c
}

func main() {
	titulos := []string{"https://www.google.com", "https://www.facebook.com", "https://www.youtube.com", "https://www.amazon.com", "https://www.wikipedia.org"}
	titulos2 := []string{"https://www.reddit.com", "https://www.twitter.com", "https://www.linkedin.com", "https://www.instagram.com", "https://www.netflix.com"}

	t1 := titulo(titulos)
	t2 := titulo(titulos2)

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
	fmt.Println("Terceiros:", <-t1, "|", <-t2)
	fmt.Println("Quartos:", <-t1, "|", <-t2)
	fmt.Println("Quintos:", <-t1, "|", <-t2)

}
