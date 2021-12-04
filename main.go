package main

import (
	"fmt"
	"time"

	greetclient "github.com/keremdokumaci/go-grpc/greet_client"
	greetserver "github.com/keremdokumaci/go-grpc/greet_server"
)

func main() {
	go greetserver.StartServer()

	fmt.Println("Waiting 5 seconds before calling Greet Service")
	time.Sleep(time.Second * 5)
	greetclient.StartClient()
}
