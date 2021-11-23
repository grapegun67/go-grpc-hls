package main

import (
        "log"
        "context"
        "fmt"
        "time"
        "google.golang.org/grpc"
        pb "test_proto/test_proto"
)

func main() {
        fmt.Println("Start Program");

        c, e := grpc.Dial("localhost:12345", grpc.WithInsecure())
        if e != nil {
                log.Fatalf("Failed to load default features: %v", e)
        }
        defer c.Close();

        fmt.Println("Start Client");

        c2 := pb.NewTestProtoClient(c)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        fmt.Println("[1]");
        r, err := c2.GetValue(ctx, &pb.FirstValue{Val1: 1234, Val2: 4321})
        if err != nil {
                log.Fatalf("Getvalue err %v", err)
        }

        fmt.Println("get value: [%d] [%d]", r.GetVal3(), r.GetVal4())
        fmt.Println("End Program");
}
