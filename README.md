# ProtoHTTP

ProtoHTTP is a simple gRPC server implementation in Go that responds with a "Hello, World!" message. This project serves as an example of how to set up a gRPC service with Go, including server reflection for tooling support.

## Features

- Simple gRPC service that responds to `SayHello` requests.
- Server reflection enabled for easy testing with tools like `grpcurl`.
- Clean and well-documented Go code.

## Getting Started

These instructions will get your copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (1.22 or higher)
- Protocol Buffer Compiler (`protoc`)
- gRPC Go Plugin
- gRPCurl (for testing)

### Installing

First, clone the repository to your local machine:

```sh
git clone https://github.com/H0llyW00dzZ/ProtoHTTP.git
cd ProtoHTTP
```

Next, install the necessary dependencies:

```sh
go mod tidy
```

### Running the server

To run the server, execute:

```sh
go run main.go
```

The server will start listening on port `50051`.

### Testing with grpcurl

You can test the `SayHello` service using `grpcurl`:

List services:

```sh
grpcurl -plaintext localhost:50051 list
```

Call the `SayHello` method:

```sh
grpcurl -plaintext -d '{}' localhost:50051 redirect.HelloService/SayHello
```

Call the `SayHelloWithField` method:

```sh
grpcurl -plaintext -d '{"name": "Gopher"}' localhost:50051 redirect.HelloService/SayHelloWithField
```

You should receive a response like:

For `SayHello`:

```json
{
  "message": "Hello, World!"
}
```

For `SayHelloWithField`:

```json
{
  "message": "Hello, Gopher!"
}
```

## Project Structure

- `main.go`: Contains the server implementation and the entry point to start the server.
- `/handler`: This directory contains the generated gRPC code from your `.proto` files.

# Acknowledgment

After I learning how Protocol Buffers work, it seems that Protocol Buffers have limitations. For example, you cannot use concurrency or manage flexibility with goroutines. It is recommended to use Protocol Buffers for projects such as real-time communication or end-to-end communication.
