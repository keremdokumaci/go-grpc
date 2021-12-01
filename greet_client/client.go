package greetclient

import (
	"context"
	"log"

	"github.com/keremdokumaci/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func StartClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // do not use WıthInsecure in prod. Its about SSL certification.
	if err != nil {
		log.Fatalf("Couldn't connect to localhost:50051 \n%v", err)
	}

	defer conn.Close() //close connection when client ends up.

	client := greetpb.NewGreetServiceClient(conn)

	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kerem",
			LastName:  "Dokumacı",
		},
	}

	response, err := client.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet RPC : %v", response.Result)
}
