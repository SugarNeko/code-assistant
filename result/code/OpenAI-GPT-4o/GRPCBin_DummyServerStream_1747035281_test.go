package grpcbin_test

import (
	"context"
	"log"
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
		FString:  "test",
		FInt32:   42,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub_test"},
		FBool:    true,
		FInt64:   1234567890,
		FBytes:   []byte("bytes_test"),
		FFloat:   3.14,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Error calling DummyServerStream: %v", err)
	}

	for i := 0; i < 10; i++ {
		res, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive message: %v", err)
		}
		if res == nil || res.FString != req.FString || res.FInt32 != req.FInt32 {
			t.Errorf("Unexpected response: got %v, want %v", res, req)
		}
	}

	log.Println("Test completed successfully")
}
