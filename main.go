package main

import (
	"log"
	"net"

	"example/gogrpc/grpcsetup"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
	grpcsetup.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *grpcsetup.Message) (*grpcsetup.Message, error) {
	log.Printf("Received message from client: %s", message.Body)
	return &grpcsetup.Message{Body: "Hello from the Server!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	grpcsetup.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
