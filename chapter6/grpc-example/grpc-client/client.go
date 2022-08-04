package main

import (
	"context"
	"log"

	"github.com/SantiagoBedoya/restfull-webservices/chapter6/grpc-example/datafiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := datafiles.NewMoneyTransactionClient(conn)
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	r, err := c.MakeTransaction(context.Background(), &datafiles.TransactionRequest{
		From:   from,
		To:     to,
		Amount: amount,
	})
	if err != nil {
		log.Fatalf("could not transact: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Confirmation)
}
