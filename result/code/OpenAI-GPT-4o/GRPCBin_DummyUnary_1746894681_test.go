package grpcbin_test

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "sub_test"},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Client response validation
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

	// Server response validation
	// Add additional checks for other fields as necessary
}
