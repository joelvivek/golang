package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I am very keen on %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9092", nil)
}