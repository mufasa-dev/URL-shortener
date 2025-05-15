package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenUrl)
	http.HandleFunc("/", handlers.RedirectHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
