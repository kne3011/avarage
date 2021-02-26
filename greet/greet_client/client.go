package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"com.grpc.tleu/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doLongGreet(c)
}

func doLongGreet(c greetpb.GreetServiceClient) {
	fmt.Println("Enter  numbers to calculate average")
	var num1 float32
	var num2 float32
	var num3 float32
	var num4 float32
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)
	fmt.Scan(&num4)

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				Number: num1,
			},
		},
		{
			Greeting: &greetpb.Greeting{
				Number: num2,
			},
			},
		{
			Greeting: &greetpb.Greeting{
				Number: num3,
			},
		},
		{
			Greeting: &greetpb.Greeting{
				Number: num4,
			},
		},
	}

	ctx := context.Background()
	stream, err := c.LongGreet(ctx)
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(2000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}
