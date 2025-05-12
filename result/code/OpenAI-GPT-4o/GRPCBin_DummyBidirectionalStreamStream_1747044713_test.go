package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Error creating stream: %v", err)
	}

	// Test data
	message := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub_test"},
		FBool:    true,
		FInt64:   456,
		FFloat:   1.23,
	}

	// Send message
	if err := stream.Send(message); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive message
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Validate response
	if resp.FString != message.FString || resp.FInt32 != message.FInt32 || resp.FEnum != message.FEnum {
		t.Errorf("Response validation failed: got %v, want %v", resp, message)
	}
}
