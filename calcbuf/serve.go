package main

import (
	"fmt"
	"net/http"
)

func serve(){

	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("/", welcome)
	http.ListenAndServe(":5000", mux)
}

func welcome(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Welcome to a calcbuf istance\n")
}
