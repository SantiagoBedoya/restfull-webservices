package main

import (
	"context"
	"io"
	"log"

	"github.com/SantiagoBedoya/restfull-webservices/chapter6/serverPush/datafiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func ReceiveStream(client datafiles.MoneyTransactionClient, request *datafiles.TransactionRequest) {
	log.Println("started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := datafiles.NewMoneyTransactionClient(conn)
	from := "1234"
	to := "5678"
	amount := float32(1250.75)
	ReceiveStream(client, &datafiles.TransactionRequest{From: from, To: to, Amount: amount})
}
