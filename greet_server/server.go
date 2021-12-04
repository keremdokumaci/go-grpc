package greetserver

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/keremdokumaci/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("Greet Unary Started ...")
	firstName := req.Greeting.GetFirstName()
	message := "Hello  " + firstName
	response := &greetpb.GreetResponse{
		Result: message,
	}

	return response, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Println("Greet Many Times Streaming Started...")
	firstName := req.GetGreeting().FirstName
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " Number: " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(time.Second * 1)
	}
	return nil
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
