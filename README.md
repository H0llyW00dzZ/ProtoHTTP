<pre>
   ______      ____             __        __  __________________ 
  / ____/___  / __ \_________  / /_____  / / / /_  __/_  __/ __ \
 / / __/ __ \/ /_/ / ___/ __ \/ __/ __ \/ /_/ / / /   / / / /_/ /
/ /_/ / /_/ / ____/ /  / /_/ / /_/ /_/ / __  / / /   / / / ____/     Example
\____/\____/_/   /_/   \____/\__/\____/_/ /_/ /_/   /_/ /_/      
                                                                 
Copyright (©️) 2024 @H0llyW00dzZ All rights reserved.
</pre>

ProtoHTTP is a simple gRPC server implementation in Go that responds with a "Hello, World!" message. This project serves as an example of how to set up a gRPC service with Go, including server reflection for tooling support.

## Features

- Simple gRPC service that responds to `SayHello` requests.
- Server reflection enabled for easy testing with tools like `grpcurl`.
- Clean and well-documented Go code.

## Getting Started

These instructions will get your copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://go.dev/dl/) (1.22 or higher)
- [Protocol Buffer Compiler](https://protobuf.dev/downloads/) (`protoc`)
- [gRPC Go Plugin](https://github.com/grpc/grpc-go)
- [gRPCurl](https://github.com/fullstorydev/grpcurl) (for testing)

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

Call the `Ping` method:

```sh
grpcurl -plaintext -d '{}' localhost:50051 redirect.PingService/Ping
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

For `Ping`:

```json
{
  "latencyMicroseconds": "101579",
  "formattedLatency": "101579 µs",
  "formattedLatencyMilliseconds": "101.58 ms"
}
```

## Project Structure

- `main.go`: Contains the server implementation and the entry point to start the server.
- `/handler`: This directory contains the generated gRPC code from your `.proto` files.

# Acknowledgment

After I learning how Protocol Buffers work, it seems that Protocol Buffers have limitations. For example, you cannot use concurrency or manage flexibility with goroutines. It is recommended to use Protocol Buffers for projects such as real-time communication or end-to-end communication.
