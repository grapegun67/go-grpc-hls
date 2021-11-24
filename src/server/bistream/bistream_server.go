package main

import (
        "log"
        "fmt"
        "net"
	"io"
        "context"
        "google.golang.org/grpc"
        pb "proto/unary"
)

type testServer struct {
        pb.UnimplementedTestProtoServer
}

func (s *testServer) GetValue(ctx context.Context, value *pb.FirstValue) (*pb.SecondValue, error) {
        log.Printf("Received: %v %v", value.GetVal1(), value.GetVal2())
        return &pb.SecondValue{Val3: value.GetVal2(), Val4: value.GetVal1()}, nil
}

func (s *testServer) GetStreamValue(value *pb.FirstValue, stream pb.TestProto_GetStreamValueServer) error {
        log.Printf("Received: %v %v", value.GetVal1(), value.GetVal2())
	for ii := 0; ii < 5; ii++ {
			if err := stream.Send(&pb.SecondValue{Val3: value.GetVal2(), Val4: value.GetVal1()}); err != nil {
				return err
		}
	}
	return nil
}

func (s *testServer) GetBiStreamValue(stream pb.TestProto_GetBiStreamValueServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Received: %v %v", in.Val1, in.Val2)

		for ii := 0; ii < 5; ii++ {
			if err := stream.Send(&pb.SecondValue{Val3: in.Val1, Val4: in.Val2}); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
        fmt.Println("Start Program");

        l, e := net.Listen("tcp", ":12345")
	if e != nil {
		log.Fatalf("failed to listen: %v", e)
	}
        defer l.Close();

        fmt.Println("Start Server");

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTestProtoServer(grpcServer, &testServer{})

        log.Printf("server listen %v", l.Addr())

	grpcServer.Serve(l)

        fmt.Println("End Program");
}
