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
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	req := &grpcbin.DummyMessage{
		FString:   "test",
		FInt32:    42,
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		FBool:     true,
		FInt64:    123456789,
		FBytes:    []byte("bytes"),
		FFloat:    1.23,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	var responses int
	for {
		_, err := stream.Recv()
		if err != nil {
			break
		}
		responses++
	}

	if responses != 10 {
		t.Errorf("Expected 10 responses, got %d", responses)
	}
}
