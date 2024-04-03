package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

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

// pingServer implements the PingServiceServer interface from the generated code.
type pingServer struct {
	handler.UnimplementedPingServiceServer
}

// Ping handles incoming PingRequests and returns a PingResponse with latency information.
func (s *pingServer) Ping(ctx context.Context, req *handler.PingRequest) (*handler.PingResponse, error) {
	// Capture the start time
	startTime := time.Now()

	// Here you would typically do some work, like database queries,
	// or any other processing your server needs to do.

	// Simulate work or a delay; you can remove this in a real-world scenario.
	time.Sleep(100 * time.Millisecond)

	// Calculate the latency
	latency := time.Since(startTime)

	// Return the response with both the raw latency in microseconds
	// and the formatted latency string
	return &handler.PingResponse{
		LatencyMicroseconds:          latency.Microseconds(),
		FormattedLatency:             fmt.Sprintf("%d µs", latency.Microseconds()),
		FormattedLatencyMilliseconds: fmt.Sprintf("%.2f ms", float64(latency.Microseconds())/1000.0),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	handler.RegisterHelloServiceServer(s, &helloServer{}) // Register the HelloService with both methods
	handler.RegisterPingServiceServer(s, &pingServer{})   // Register the PingService with the grpc server

	// Enable server reflection
	reflection.Register(s)

	// Start serving in a goroutine to allow shutdown to proceed in parallel
	go startServer(s, lis)

	// Wait for the server to shut down gracefully when an OS signal is received
	waitForShutdown(s)
}

// Start the gRPC server and listen for incoming connections
func startServer(s *grpc.Server, lis net.Listener) {
	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil && err != grpc.ErrServerStopped {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Wait for interrupt signal to gracefully shutdown the server
func waitForShutdown(s *grpc.Server) {
	// Create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	// Create a channel to receive a signal when server shutdown is complete
	done := make(chan bool, 1)

	// Register to receive specific signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine that will listen for signals
	go func() {
		sig := <-sigs
		log.Printf("Received signal: %s", sig)

		// Perform server shutdown
		log.Println("Shutting down server...")
		s.GracefulStop() // Gracefully stop the server
		log.Println("Server has been shut down.")

		// Notify the main goroutine that we're done
		done <- true
	}()

	// Wait for shutdown to complete
	<-done
	log.Println("Graceful shutdown completed")
}
