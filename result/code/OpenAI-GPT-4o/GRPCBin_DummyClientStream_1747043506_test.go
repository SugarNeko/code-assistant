package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		err = stream.Send(&grpcbin.DummyMessage{
			FString:  "test",
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FInt32:   int32(i),
			FInt64:   int64(i),
			FBool:    true,
			FFloat:   float32(i),
			FBytes:   []byte("bytes"),
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
			FStrings: []string{"str1", "str2"},
		})
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	if resp.GetFString() != "test" {
		t.Errorf("Expected FString to be 'test', got %v", resp.GetFString())
	}

	if resp.GetFEnum() != grpcbin.DummyMessage_ENUM_1 {
		t.Errorf("Expected FEnum to be ENUM_1, got %v", resp.GetFEnum())
	}
}
