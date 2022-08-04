package main

import (
	"encoding/json"
	"fmt"

	"github.com/SantiagoBedoya/restfull-webservices/chapter6/protofiles"
)

func main() {
	p := &protofiles.Person{
		Id:    1234,
		Name:  "Roger F",
		Email: "rf@google.com",
		Phones: []*protofiles.Person_PhoneNumber{
			{Number: "555-4321", Type: protofiles.Person_HOME},
		},
	}

	body, _ := json.Marshal(p)
	fmt.Println(string(body))

	// p1 := &protofiles.Person{}
	// body, _ := proto.Marshal(p)
	// _ = proto.Unmarshal(body, p1)
	// fmt.Println("Original struct loaded from proto file: ", p)
	// fmt.Println("Marshaled proto data: ", body)
	// fmt.Println("Unmarshaled struct: ", p1)
}
