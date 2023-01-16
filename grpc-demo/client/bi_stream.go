package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) error {
	log.Printf("Bidirectional streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		return nil
	}
	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional stream finished")
	return nil
}
