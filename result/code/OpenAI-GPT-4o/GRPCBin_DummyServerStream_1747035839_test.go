package grpcbin

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
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   32,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FBool:    true,
		FInt64:   64,
		FBytes:   []byte("bytes"),
		FFloat:   3.14,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	responseCount := 0
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}

		responseCount++

		if msg.FString != req.FString {
			t.Errorf("Expected FString: %v, got: %v", req.FString, msg.FString)
		}
	}

	if responseCount != 10 {
		t.Errorf("Expected 10 messages, got: %v", responseCount)
	}
}
