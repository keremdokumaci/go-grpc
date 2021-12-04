package greetclient

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	log.Println("Server streaming client..")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kerem",
			LastName:  "Dokumacı",
		},
	}

	resStream, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal("Error while calling greet many times rpc :%v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF { //reached the end of stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream :%v ", err)
		}

		log.Printf("Response from GreetManyTimesRPC : %v ", msg.GetResult())
	}

	log.Println("Client streaming.... (LongGreet)")
	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling longgreet(client streaming)... \n%v", err)
	}

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kerem",
				LastName:  "Dokumacı",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kerem2",
				LastName:  "Dokumacı2",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kerem3",
				LastName:  "Dokumacı3",
			},
		},
	}

	for _, req := range requests {
		log.Printf("Sending req : %v", req)
		stream.Send(req)
		time.Sleep(time.Millisecond * 200)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet %v", err)
	}

	fmt.Println("LongGreet response : %v", res)
}
