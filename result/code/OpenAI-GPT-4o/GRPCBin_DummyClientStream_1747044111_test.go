package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	// Set timeout for connection
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to start stream: %v", err)
	}

	// Send test messages
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString: "test",
			FInt32:  int32(i),
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	// Close the stream and receive response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate response
	if resp.FString != "test" {
		t.Errorf("Unexpected response FString: got %v, want test", resp.FString)
	}
	if resp.FInt32 != 9 {
		t.Errorf("Unexpected response FInt32: got %v, want 9", resp.FInt32)
	}
}
