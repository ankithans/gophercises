package main

import (
	"fmt"
	"net/http"

	"urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /git-ankit
  url: https://github.com/ankithans
- path: /ankithans
  url: https://ankithans.github.io
`
	yamlHandler, err := urlshort.YamlHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello Ankit!")
}
