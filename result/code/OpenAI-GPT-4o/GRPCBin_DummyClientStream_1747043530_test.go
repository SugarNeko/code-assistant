package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Send 10 messages
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.DummyMessage{
			FString: "test",
			FInt32:  int32(i),
		}); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	// Close the stream
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive reply: %v", err)
	}

	// Validate client response
	if reply.GetFString() != "test" {
		t.Errorf("Expected FString 'test', got '%s'", reply.GetFString())
	}

	if reply.GetFInt32() != 9 {
		t.Errorf("Expected FInt32 9, got %d", reply.GetFInt32())
	}

	// Additional validation can be added here
}
