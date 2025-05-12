package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

const address = "grpcb.in:9000"

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to open stream: %v", err)
	}

	// Send a test message
	testMessage := &pb.DummyMessage{
		FString:  "test",
		FStrings: []string{"test1", "test2"},
		FInt32:   123,
		FBool:    true,
	}
	if err := stream.Send(testMessage); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive a response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Validate the response
	if resp.FString != testMessage.FString || len(resp.FStrings) != len(testMessage.FStrings) || resp.FInt32 != testMessage.FInt32 || resp.FBool != testMessage.FBool {
		t.Errorf("Response did not match request. Got: %v, want: %v", resp, testMessage)
	}
}
