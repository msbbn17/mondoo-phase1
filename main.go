package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Mondoo Engineer!")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil) // nosemgrep
	if err != nil {
		log.Fatal("ListenAndServe-Fehler: ", err)
	}
}
