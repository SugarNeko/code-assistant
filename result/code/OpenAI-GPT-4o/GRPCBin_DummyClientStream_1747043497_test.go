package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	messages := []*grpcbin.DummyMessage{
		{FString: "test1"},
		{FString: "test2"},
		// Add more messages for a complete test
	}

	for _, msg := range messages {
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	expected := messages[len(messages)-1]
	if reply.FString != expected.FString {
		t.Errorf("Unexpected response. Got: %v, want: %v", reply.FString, expected.FString)
	}
}
