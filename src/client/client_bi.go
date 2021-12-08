package main

import (
	"sync"
        "log"
        "context"
        "fmt"
        "time"
        "google.golang.org/grpc"
        pb "proto/unary"
)

func main() {
        fmt.Println("Start Program");

	/* go routine sync를 위한 변수 선언	*/
	var wait sync.WaitGroup
	var num int = 0;
	wait.Add(2)

        c, e := grpc.Dial("localhost:12345", grpc.WithInsecure())
        if e != nil {
                log.Fatalf("Failed to load default features: %v", e)
        }
        defer c.Close();

        fmt.Println("Start Client");

        c2 := pb.NewTestProtoClient(c)
        ctx := context.Background()

	stream, err := c2.GetBiStreamValue(ctx)
	if err != nil {
		log.Fatalf("stream error: %v", err)
	}

	/* Recv를 위한 고루틴	*/
	go func() {

		log.Printf("==================<Start RECV GO ROUTINE>==================")

		for {
			in, err := stream.Recv()
			/* 의도적으로 go routine 을 종료시키기 위한 임시 방편	*/
			if num > 2 {
				fmt.Println("END I/O!!")
				defer wait.Done()

				break
			}

			log.Printf("------[Recv Message]------")
			log.Printf("Got message: val3[%d], val4[%d]", in.Val3, in.Val4)

			if err != nil {
				log.Fatalf("stream.Recv() ERROR: %v", err)
			}

			log.Printf("nil: [%v]", err)
		}

	}()

	/* Bistream 시작 	*/
	go func() {

		log.Printf("==================<Start SEND GO ROUTINE>==================")

		for {
			log.Printf("------[Send Message]------")

			j := int32(22)

			for i := int32(34); i < 37; i++ {
				if err := stream.Send(&pb.FirstValue{Val1: i, Val2:j}); err != nil {
					log.Fatalf("Send ERROR!! %v", err)
				}
			}

			num++
			log.Printf("num: [%d]", num)
			if num > 2 {
				defer wait.Done()

				break
			}

			log.Printf("------[Sleep]------")
			time.Sleep(4000 * time.Millisecond)
		}

	}()

	wait.Wait()

        fmt.Println("End Program");
}
/************************************************************************************
[정리]
1. send와 recv는 언제 할지 모르므로 절차적 프로그램으로 처리할 수 없다. 따라서 고루틴을 사용해서
각자 정해진 역할대로 Recv하고 Send 한 뒤 기능을 수행하도록 만들어야 한다.


*************************************************************************************/


/************************************************************************************
[Reference]
1. stream.Recv()에 대한 End of FILE 처리. 작동하지는 않음. 참고만 해볼 것
	source)

import	"io"

		in, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("END I/O!!")
			defer wait.Done()

			break
		}

*************************************************************************************/
