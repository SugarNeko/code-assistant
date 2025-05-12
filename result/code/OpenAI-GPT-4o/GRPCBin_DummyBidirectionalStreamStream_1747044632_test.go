package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Error creating stream: %v", err)
	}

	request := &grpcbin.DummyMessage{
		FString: "Test String",
		FInt32:  123,
	}

	if err := stream.Send(request); err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	
	response, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	if response.FString != request.FString || response.FInt32 != request.FInt32 {
		t.Errorf("Unexpected response: got %v, want %v", response, request)
	}

	log.Printf("Client received: %v", response)
}
