package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// The main function begins with a call to http.HandleFunc,
	// which tells the http package to handle all requests to the web root ("/") with handler.
	http.HandleFunc("/", handler)

	// It then calls http.ListenAndServe,
	// specifying that it should listen on port 8080 on any interface (":8080").
	// (Don't worry about its second parameter, nil, for now.)
	// This function will block until the program is terminated.
	log.Fatal(http.ListenAndServe(":8080", nil))
}