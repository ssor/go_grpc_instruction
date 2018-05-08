package main

import (
    pb "demo/hello_grpc"
    "google.golang.org/grpc"
    "log"
    "os"
    "context"
    "time"
)

var (
    grpcTarget = "localhost:5000"
)

func main() {
    conn, err := grpc.Dial(grpcTarget, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to dial grpc server: %s", err)
    }
    var name string
    if len(os.Args) > 1 {
        name = os.Args[1]
    } else {
        name = "no-name"
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    client := pb.NewGreeterClient(conn)
    res, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("failed to say hello to grpc server: %s", err)
    }
    log.Println(name, " back with: ", res.Message)
}
