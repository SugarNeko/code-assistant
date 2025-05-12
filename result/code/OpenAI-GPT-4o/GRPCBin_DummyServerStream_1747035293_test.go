package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyServerStream_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyServerStream(context.Background(), &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  123,
		// add more fields based on your requirement
	})
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	expectedCount := 10
	var receivedCount int
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		receivedCount++

		if resp.FString != "test" {
			t.Errorf("Expected FString 'test', got '%s'", resp.FString)
		}
	}

	if receivedCount != expectedCount {
		t.Errorf("Expected %d responses, got %d", expectedCount, receivedCount)
	}
}
