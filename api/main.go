package main

import (
	"log"
	"net/http"

	"now-go-mod-child/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Child)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
