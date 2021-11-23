package main

import (
        "log"
        "fmt"
        "net"
        "context"
        "google.golang.org/grpc"
        "google.golang.org/grpc/reflection"
        pb "proto/unary"
)

type testServer struct {
        pb.UnimplementedTestProtoServer
}

func (s *testServer) GetValue(ctx context.Context, value *pb.FirstValue) (*pb.SecondValue, error) {
        log.Printf("Received: %v %v", value.GetVal1(), value.GetVal2())
        return &pb.SecondValue{Val3: value.GetVal2(), Val4: value.GetVal1()}, nil
}

func main() {
        fmt.Println("Start Program");

        l, e := net.Listen("tcp", ":12345")
        if e != nil {
                log.Fatalf("Failed to load default features: %v", e)
        }
        defer l.Close();

        fmt.Println("Start Server");
        s := grpc.NewServer()
        pb.RegisterTestProtoServer(s, &testServer{})
        log.Printf("server listen %v", l.Addr())

        reflection.Register(s)
        if e = s.Serve(l); e != nil {
                log.Fatalf("Failed to Server()")
        }

        fmt.Println("End Program");
}
