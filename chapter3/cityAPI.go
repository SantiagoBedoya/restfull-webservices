package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

type city struct {
	Name string
	Area uint64
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(http.StatusText(http.StatusUnsupportedMediaType)))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&tempCity); err != nil {
			panic(err)
		}
		defer r.Body.Close()
		fmt.Printf("Got %s city with area of %d sq miles\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(mainLogicHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8080", nil)
}
