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
		FString:   "test",
		FInt32:    123,
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		FInt64:    456,
		FBytes:    []byte("bytes"),
		FFloat:    1.23,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			closeErr := stream.CloseSend()
			if closeErr != nil {
				t.Errorf("Error closing stream: %v", closeErr)
			}
			break
		}
		if resp.FString != req.FString {
			t.Errorf("Expected FString: %s, got: %s", req.FString, resp.FString)
		}
		count++
	}

	if count != 10 {
		t.Errorf("Expected 10 messages, received %d", count)
	}
}
