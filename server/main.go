package main

import (
    "context"
    pb "demo/hello_grpc"
    "log"
    "net"

    "google.golang.org/grpc"
)

var (
    port = ":5000"
)

func main() {
    listener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen on %s, error: %v", port, err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterGreeterServer(grpcServer, &server{})
    log.Println("listen on ", port)
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("failed to serve grpc: %v", err)
    }
}

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    log.Println(in.Name, " comes in")
    res := &pb.HelloResponse{
        Message: "Hello, " + in.Name,
    }
    return res, nil
}
