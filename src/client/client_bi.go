package main

import (
	"io"
        "log"
        "context"
        "fmt"
        "time"
        "google.golang.org/grpc"
        pb "proto/unary"
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
        ctx, cancel := context.WithTimeout(context.Background(), 1000 * time.Second)
        defer cancel()


	stream, err := c2.GetBiStreamValue(ctx)
	if err != nil {
		log.Fatalf("stream error: %v", err)
	}

	/* Bistream 시작 	*/
	for {
		j := int32(22)
		for i := int32(34); i < 37; i++ {
			if err := stream.Send(&pb.FirstValue{Val1: i, Val2:j}); err != nil {
				log.Fatalf("Send ERROR!! %v", err)
			}
		}

		for {
			in, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("No I/O!!")
				break
			}

			log.Printf("Got message: val3[%d], val4[%d]", in.Val3, in.Val4)

			if err != nil {
				log.Fatalf("stream.Recv() ERROR: %v", err)
			} else {
				/* recv와 send가 고루틴으로  처리되어야하나? 이게 잘 안되네	*/
				break
			}
		}

		time.Sleep(4000 * time.Millisecond)
	}

        fmt.Println("End Program");
}
