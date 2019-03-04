package gistdirect

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
