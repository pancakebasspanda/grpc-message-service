package main

import (
	"context"
	"flag"
	"io"
	"log"
	"pancakebasspanda/grpc-message-service/client"
	"pancakebasspanda/grpc-message-service/protos"
)

var (
	connectionID string
)

func init() {
	flag.StringVar(&connectionID, "connectionID", "", "unique identifier for the connection")
}

func main() {

	ctx := context.Background()

	messageService, err := client.New(ctx)
	if err != nil {
		panic("error connecting to message-service")
	}

	stream, err := messageService.Connection(ctx)
	waitc := make(chan struct{})

	// receiving
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s", in)
		}
	}()

	// sending
	if err := stream.Send(&protos.ConnectionRequest{
		ConnectionId: connectionID,
	}); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}

	stream.CloseSend()
	<-waitc

}
