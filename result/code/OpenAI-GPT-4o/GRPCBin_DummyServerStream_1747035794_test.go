package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   42,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FBool:    true,
		FInt64:   64,
		FBytes:   []byte("bytes"),
		FFloat:   3.14,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		// Validate client response
		if resp.FString != req.FString {
			t.Errorf("Expected FString %v, got %v", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32*10 {
			t.Errorf("Expected FInt32 %v, got %v", req.FInt32*10, resp.FInt32)
		}
		count++
	}

	// Test server response validation
	if count != 10 {
		t.Errorf("Expected 10 responses, got %d", count)
	}
}
