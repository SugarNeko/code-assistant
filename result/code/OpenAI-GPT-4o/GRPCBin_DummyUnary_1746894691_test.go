package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test_string",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "sub_string"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString: %s, got: %s", req.FString, resp.FString)
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32: %d, got: %d", req.FInt32, resp.FInt32)
	}

	if resp.FEnum != req.FEnum {
		t.Errorf("Expected FEnum: %v, got: %v", req.FEnum, resp.FEnum)
	}

	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("Expected FSub.FString: %s, got: %s", req.FSub.FString, resp.FSub.FString)
	}
}
