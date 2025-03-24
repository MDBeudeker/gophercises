package main

import (
	"fmt"
	"net/http"
	"example.com/02/urlshort"
)

func main() {

	mux := defaultMux()

	pathsToUrls:= map[string]string{
		"/reddit": "https://old.reddit.com",
		"/github": "https://github.com",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the Maphandler as the fallback
	yaml := `
- path: /reddit
  url: https://old.reddit.com
- path: /thebox
  url: https://themetalbox.com/
`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	port :=":8080"

	fmt.Print("starting server at localhost",port)
	http.ListenAndServe(port, yamlHandler)
	
}



func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// TODO

// invoke the YAML handling
// write flag that loads the YAML file
// invode JSON handling

// error handling
// read from database like BoltDB

