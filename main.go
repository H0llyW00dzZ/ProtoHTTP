package main

import (
	"context"
	"log"
	"net"

	"github.com/H0llyW00dzZ/ProtoHTTP/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// helloServer implements the HelloServiceServer interface from the generated code.
type helloServer struct {
	handler.UnimplementedHelloServiceServer
}

// SayHello handles incoming HelloRequests and returns a HelloResponse.
func (s *helloServer) SayHello(ctx context.Context, req *handler.HelloRequest) (*handler.HelloResponse, error) {
	return &handler.HelloResponse{
		Result: &handler.HelloResponse_Message{Message: "Hello, World!"},
	}, nil
}

func (s *helloServer) SayHelloWithField(ctx context.Context, req *handler.HelloRequestWithField) (*handler.HelloResponse, error) {
	if req.Name == "" {
		st := status.New(codes.InvalidArgument, "Name field is required")
		return &handler.HelloResponse{
			Result: &handler.HelloResponse_Error{Error: st.Proto()},
		}, nil
	}

	// If a name is provided, return a successful response.
	return &handler.HelloResponse{
		Result: &handler.HelloResponse_Message{Message: "Hello, " + req.Name + "!"},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	handler.RegisterHelloServiceServer(s, &helloServer{}) // Register the HelloService with both methods

	// Enable server reflection
	reflection.Register(s)

	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
