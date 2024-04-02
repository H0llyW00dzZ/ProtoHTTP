package main

import (
	"context"
	"log"
	"net"

	"github.com/H0llyW00dzZ/ProtoHTTP/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// helloServer implements the HelloServiceServer interface from the generated code.
type helloServer struct {
	handler.UnimplementedHelloServiceServer
}

// SayHello handles incoming HelloRequests and returns a HelloResponse.
func (s *helloServer) SayHello(ctx context.Context, req *handler.HelloRequest) (*handler.HelloResponse, error) {
	return &handler.HelloResponse{Message: "Hello, World!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	handler.RegisterHelloServiceServer(s, &helloServer{}) // Register the HelloService

	// Enable server reflection
	reflection.Register(s)

	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
