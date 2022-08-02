package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got paramter id: %s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got paramter category: %s\n", queryParams["category"][0])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
