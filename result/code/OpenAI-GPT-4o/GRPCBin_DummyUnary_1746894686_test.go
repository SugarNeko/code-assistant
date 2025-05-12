package grpcbin_test

import (
	"context"
	"testing"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
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
		FSub:    &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FBool:   true,
		FInt64:  12345,
		FFloat:  1.23,
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != "test" || resp.FInt32 != 123 || resp.FEnum != grpcbin.DummyMessage_ENUM_1 || !resp.FBool || resp.FInt64 != 12345 || resp.FFloat != 1.23 {
		t.Errorf("Unexpected response: %+v", resp)
	}

	if resp.FSub.FString != "subtest" {
		t.Errorf("Unexpected sub message: %+v", resp.FSub)
	}
}
