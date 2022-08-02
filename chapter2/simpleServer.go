package main

import (
	"io"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

// func main() {
// 	http.HandleFunc("/hello", handleHello)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
