package titulo

import (
	"io"
	"net/http"
	"regexp"
)

// Titulo function obtem o titulo de uma pagina web
func Titulo(urls []string) <-chan string {
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

func TituloStream(urls ...string) <-chan string {
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
