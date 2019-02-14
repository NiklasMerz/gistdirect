package main

import (
	"fmt"
	"net/http"

	"github.com/jaytaylor/go-hostsfile"
)

func main() {
	res, err := hostsfile.ReverseLookup("127.0.0.1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v maps 127.0.0.1 to the following names: %v", hostsfile.HostsPath, res)
}

// Handler is the serverless entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	res, err := hostsfile.ReverseLookup("127.0.0.1")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v maps 127.0.0.1 to the following names: %v", hostsfile.HostsPath, res)
}
