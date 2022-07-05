package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Welcome)
	http.ListenAndServe(":8080", nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
