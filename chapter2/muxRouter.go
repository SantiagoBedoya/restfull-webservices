package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")

// 	srv := &http.Server{
// 		Addr:         ":8080",
// 		Handler:      r,
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}
// 	log.Fatal(srv.ListenAndServe())
// }
