package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "subtest"},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	var count int
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err.Error() != "EOF" {
				t.Fatalf("Failed to receive stream: %v", err)
			}
			break
		}

		// Validate response
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
			t.Errorf("Expected FSub FString %v, got %v", req.FSub.FString, resp.FSub.FString)
		}

		count++
	}

	// Validate the number of responses
	if count != 10 {
		t.Errorf("Expected 10 responses, got %v", count)
	}
}
