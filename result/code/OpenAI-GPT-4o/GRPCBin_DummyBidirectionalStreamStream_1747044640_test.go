package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"your-module-path/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	// Set up the connection with a timeout
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Example request sending
	req := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	// Receive response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate the response
	if resp.FString != req.FString || resp.FInt32 != req.FInt32 || resp.FEnum != req.FEnum {
		t.Error("Response does not match the request")
	}

	// Log success
	log.Println("TestDummyBidirectionalStreamStream passed")
}
