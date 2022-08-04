package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/SantiagoBedoya/restfull-webservices/chapter6/serverPush/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port      = ":8080"
	noOfSteps = 3
)

type server struct{}

func (s *server) MakeTransaction(in *datafiles.TransactionRequest, stream datafiles.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money transfer...")
	log.Printf("Amount: $%f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)
	for i := 0; i < noOfSteps; i++ {
		time.Sleep(time.Second * 2)
		if err := stream.Send(&datafiles.TransactionResponse{
			Status:      "Good",
			Step:        int32(i),
			Description: fmt.Sprintf("Description of step %d", int32(i)),
		}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}
	log.Printf("Sucessfully trasfered amount $%v from %v to %v", in.Amount, in.From, in.To)
	return nil
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
		log.Fatalf("Failed to serve: %v", err)
	}
}
