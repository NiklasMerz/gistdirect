package main

import (
	"log"
	"net/http"

	"github.com/NiklasMerz/gistdirect"
)

// Main function for use as binary
func main() {
	http.HandleFunc("/", gistdirect.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
