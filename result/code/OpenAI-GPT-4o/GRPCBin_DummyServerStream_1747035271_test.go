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
	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub_test"},
		FBool:    true,
		FInt64:   123456789,
		FBytes:   []byte("bytes_test"),
		FFloat:   123.45,
	}
	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive data: %v", err)
		}
		if resp == nil {
			t.Fatalf("Expected a valid response, got nil")
		}

		// Validate response
		if resp.FString != req.FString {
			t.Errorf("Expected FString %v, got %v", req.FString, resp.FString)
		}
		// Add more validation checks as necessary
	}
}
