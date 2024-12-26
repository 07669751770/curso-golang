package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple HTTP server
	http.HandleFunc("/usuarios/", UsuarioHandler)

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
