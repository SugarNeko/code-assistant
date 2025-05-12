package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream(t *testing.T) {
	// Setup connection with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	request := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   42,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub_test"},
		FInt64:   1234567890,
		FBytes:   []byte("byte_test"),
		FFloat:   3.14,
	}

	stream, err := client.DummyServerStream(ctx, request)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	count := 0
	for {
		response, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			t.Errorf("Failed to receive stream: %v", err)
		}

		if response == nil {
			t.Errorf("Received nil response")
		} else {
			// Validate response content
			if response.FString != request.FString {
				t.Errorf("Expected FString: %v, received: %v", request.FString, response.FString)
			}
			if response.FInt32 != request.FInt32*10 {
				t.Errorf("Expected FInt32: %v, received: %v", request.FInt32*10, response.FInt32)
			}
		}

		count++
	}

	if count != 10 {
		t.Errorf("Expected to receive 10 responses, got %d", count)
	}
}
