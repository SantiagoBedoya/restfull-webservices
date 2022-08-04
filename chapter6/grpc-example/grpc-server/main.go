package main

import (
	"context"
	"log"
	"net"

	"github.com/SantiagoBedoya/restfull-webservices/chapter6/grpc-example/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) MakeTransaction(ctx context.Context, in *datafiles.TransactionRequest) (*datafiles.TransactionResponse, error) {
	log.Printf("Got request for money transfer...")
	log.Printf("Amount: %f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)
	return &datafiles.TransactionResponse{Confirmation: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	datafiles.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
