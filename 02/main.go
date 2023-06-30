package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello world!")
	})

	port :=":8080"

	fmt.Print("starting server at localhost",port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// TODO

// write code to serve a server that listens on localhost:8080

// write code to show two urls that forward to a different website

// invoke the YAML handling
// write flag that loads the YAML file
// invode JSON handling

// error handling
// read from database like BoltDB

