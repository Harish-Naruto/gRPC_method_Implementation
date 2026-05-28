# gRPC Method Implementation (Go)

This repository demonstrates all four primary gRPC communication patterns in Go:

- Unary RPC
- Server streaming RPC
- Client streaming RPC
- Bidirectional streaming RPC

It is structured as a learning/reference project with separate client and server examples for each pattern.

## What this project demonstrates

This project uses one shared protobuf contract (`proto/user.proto`) and implements it in four different RPC styles so you can compare behavior and request/response flow side-by-side.

## Tech Stack

- Go `1.25`
- gRPC (`google.golang.org/grpc`)
- Protocol Buffers (`google.golang.org/protobuf`)
- Buf (`buf.yaml`, `buf.gen.yaml`) for protobuf generation workflow

## Repository Structure

```text
.
├── proto/                       # Protobuf service and message definitions
├── gen/go/proto/                # Generated Go code from protobuf
├── unary-stream/                # Unary RPC example (client/server)
├── server-stream/               # Server streaming RPC example (client/server)
├── client-stream/               # Client streaming RPC example (client/server)
└── bidirectional-stream/        # Bidirectional streaming RPC example (client/server)
```

## Prerequisites

- Go `1.25+`
- Protocol Buffers compiler (`protoc`) if regenerating code
- Buf CLI (optional, for generation using Buf config)

## Setup

Install dependencies:

```bash
go mod download
```

## Running Examples

Each example has its own server and client under the corresponding folder.  
Start a server first, then run its client in a separate terminal.

### 1) Unary RPC

```bash
go run ./unary-stream/server
go run ./unary-stream/client
```

Client sends a single request and receives a single response.

### 2) Server Streaming RPC

```bash
go run ./server-stream/server
go run ./server-stream/client
```

Client sends one request and receives multiple streamed responses.

### 3) Client Streaming RPC

```bash
go run ./client-stream/server
go run ./client-stream/client
```

Client sends multiple streamed requests and receives one response at the end.

### 4) Bidirectional Streaming RPC

```bash
go run ./bidirectional-stream/server
go run ./bidirectional-stream/client
```

Client and server both send streams simultaneously over one connection.

## Protocol Definition

The core API contract is defined in:

- `proto/user.proto`

## Regenerating Protobuf Stubs

If you update protobuf definitions, regenerate code with Buf:

```bash
buf generate
```

Generation configuration files:

- `buf.yaml`
- `buf.gen.yaml`

## Notes

- Default examples use port `9000`.
- Run only one server at a time on that port.
