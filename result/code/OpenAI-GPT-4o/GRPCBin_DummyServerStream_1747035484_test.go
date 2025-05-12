package grpcbin_test

import (
	"context"
	"testing"
	"time"

	grpcbin "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test_string",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "sub_string"},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	expectedCount := 10
	receivedCount := 0

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		receivedCount++

		// Validate response values as necessary
		if resp.FString != req.FString {
			t.Errorf("Expected FString %v, got %v", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Expected FInt32 %v, got %v", req.FInt32, resp.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Expected FEnum %v, got %v", req.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("Expected FSub.FString %v, got %v", req.FSub.FString, resp.FSub.FString)
		}
	}

	if receivedCount != expectedCount {
		t.Errorf("Expected to receive %d messages, but received %d", expectedCount, receivedCount)
	}
}
