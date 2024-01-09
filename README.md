# gRPC Server Example

Welcome to the gRPC Server Example! This example demonstrates how to use gRPC to create a simple server application using Protoc, gRPC, and the Evans CLI.

## Technologies Used

- [Protoc](https://grpc.io/docs/languages/go/quickstart/): Protocol Buffers compiler used to generate gRPC code.
- [gRPC](https://grpc.io/): A high-performance, open-source framework for building remote procedure call (RPC) APIs.
- [Evans CLI](https://github.com/ktr0731/evans): Interactive REPL for gRPC services.

## Example Overview

This example illustrates the process of creating a gRPC server that can receive requests from a client (Evans) and send responses back. Here's a step-by-step guide:

### How it Works

1. **Proto File Creation**: Design a proto file with a service definition.
2. **Code Generation**: Generate gRPC code from the proto file using Protoc.
   - The generated code will be located in the `internal/pb` folder.
   - Notable file: `course_category_grpc.pb.go` containing gRPC code with structs, methods, and interfaces.
3. **Service Implementation**: Implement services based on the generated interfaces in `internal/service`.
4. **Server Creation**: Develop the server, located in `cmd/grpcServer`.
5. **Run the Server**: Execute the server application.
6. **Run Evans CLI**: Open Evans CLI and set the package and service.
7. **Call Service**: Use Evans CLI to call the service and send a request.
8. **Receive Response**: Observe the response from the server.

### How to Run

1. **Run the Server**:
   ```bash
   go run cmd/grpcServer/main.go
   ```
2. **Run Evans CLI**:
   ```bash
   evans -r repl
   ```
   - Set the package:
     ```evans
     $ package pb
     ```
   - Set the service:
     ```evans
     $ service CategoryService
     ```
   - Call the service (any service):
     ```evans
     $ call <AnyServiceHere>
     ```

### Example

```bash
$ evans -r repl
```

```evans
$ package pb
```

```evans
$ service CategoryService
```

```evans
$ call ListCategories
```

Feel free to explore and modify this example to suit your specific use case!