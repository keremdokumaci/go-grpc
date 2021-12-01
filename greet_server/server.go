package greetserver

import (
	"context"
	"log"
	"net"

	"github.com/keremdokumaci/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.Greeting.GetFirstName()
	message := "Hello  " + firstName
	response := &greetpb.GreetResponse{
		Result: message,
	}

	return response, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	log.Println("Server started ...")

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve %v", err)
	}
}
