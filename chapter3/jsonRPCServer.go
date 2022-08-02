package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	rpcJSON "github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil {
		log.Println("error: ", readerr)
		os.Exit(1)
	}
	marshallerr := json.Unmarshal(raw, &books)
	if marshallerr != nil {
		log.Println("error: ", marshallerr)
		os.Exit(1)
	}
	for _, book := range books {
		if book.ID == args.ID {
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(rpcJSON.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
