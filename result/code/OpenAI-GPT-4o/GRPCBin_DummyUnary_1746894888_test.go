package grpcbin_test

import (
	"context"
	"testing"
	"time"

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
		FString:  "test",
		FInt32:   123,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FBool:    true,
		FInt64:   123456789,
		FBytes:   []byte("bytes"),
		FFloat:   3.14,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("Failed to call DummyUnary: %v", err)
	}

	if res == nil || res.FString != "test" || res.FInt32 != 123 || res.FBool != true {
		t.Errorf("Response validation failed: got %v, want %v", res, req)
	}
}
