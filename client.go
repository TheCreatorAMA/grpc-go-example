package main

import (
	"log"

	"example/gogrpc/grpcsetup"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := grpcsetup.NewChatServiceClient(conn)

	message := grpcsetup.Message{Body: "Hello from the client!"}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
