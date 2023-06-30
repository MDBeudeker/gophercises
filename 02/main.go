package main

import (
	"fmt"
	"log"
	"net/http"
	"example.com/02/urlshort"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello world!")
	})

	port :=":8080"

	yaml := `
- path: /brave
  url: https://github.com/brave
- path /thebox
  url: https://themetalbox.com/
`

	fmt.Print (yaml)
	fmt.Print("starting server at localhost",port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// function that takes a input path and serves the upstream
func redirect(path string, upstream string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request){
		http.Redirect(w,r, upstream,http.StatusSeeOther)
	})	
}

// TODO

// invoke the YAML handling
// write flag that loads the YAML file
// invode JSON handling

// error handling
// read from database like BoltDB

