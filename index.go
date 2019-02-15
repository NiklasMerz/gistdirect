package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Handler is the serverless entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	url, ok := os.LookupEnv("GIST_URL")
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "GIST_URL environment variable required but not set")
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Cannot get gist %v", err.Error())
		return
	}
	defer resp.Body.Close()

	p, err := ioutil.ReadAll(resp.Body)
	hostmap := Parse(p)
	path := strings.Replace(r.URL.Path, "/", "", -1)
	if val, ok := hostmap[path]; ok {
		http.Redirect(w, r, val, http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found %v", path)
	}
}

// Parse takes in hosts file content and returns a map of parsed results. From https://github.com/jaytaylor/go-hostsfile
func Parse(hostsFileContent []byte) map[string]string {
	hostsMap := map[string]string{}
	for _, line := range strings.Split(strings.Trim(string(hostsFileContent), " \t\r\n"), "\n") {
		line = strings.Replace(strings.Trim(line, " \t"), "\t", " ", -1)
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}
		pieces := strings.SplitN(line, " ", 2)
		if len(pieces) > 1 && len(pieces[0]) > 0 {
			hostsMap[pieces[0]] = pieces[1]
		}
	}
	return hostsMap
}
