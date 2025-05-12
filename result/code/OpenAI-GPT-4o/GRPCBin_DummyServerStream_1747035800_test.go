package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	// DialOptions with a 15-second timeout
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Create a dummy request
	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FBool:    true,
		FInt64:   456,
		FBytes:   []byte("byte_test"),
		FFloat:   3.14,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Call the DummyServerStream method
	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("Error calling DummyServerStream: %v", err)
	}

	// Validate server response
	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Error receiving stream: %v", err)
		}

		// Validate response content
		if resp.FString != req.FString {
			t.Errorf("Expected FString: %v, got: %v", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Expected FInt32: %v, got: %v", req.FInt32, resp.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Expected FEnum: %v, got: %v", req.FEnum, resp.FEnum)
		}
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("Expected FSub FString: %v, got: %v", req.FSub.FString, resp.FSub.FString)
		}
		if resp.FBool != req.FBool {
			t.Errorf("Expected FBool: %v, got: %v", req.FBool, resp.FBool)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("Expected FInt64: %v, got: %v", req.FInt64, resp.FInt64)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("Expected FBytes: %v, got: %v", req.FBytes, resp.FBytes)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("Expected FFloat: %v, got: %v", req.FFloat, resp.FFloat)
		}
	}
}
