package main

import (
        "flag"
        "log"
        "fmt"
        "net/http"
        "context"
        "google.golang.org/grpc"
        "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
        gw "proto/unary"	/*	GOROOT/proto/gw		*/
)

var (
        grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:12345", "gRPC server endpoint")
)

func main() {
        fmt.Println("Start Program");
        ctx := context.Background()
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()

        mux := runtime.NewServeMux()
        opts := []grpc.DialOption{grpc.WithInsecure()}
        err := gw.RegisterTestProtoHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
        if err != nil {
                log.Fatalf("Failed to load default features: %v", err)
        }

        fmt.Println("Listening client...")
        http.ListenAndServe(":8080", mux)
        if err != nil {
                log.Fatalf("Failed to listen: %v", err)
        }
        fmt.Println("End Program");
}
