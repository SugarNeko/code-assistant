package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub"},
		FInt64:   123456789,
		FFloat:   3.14,
		FBytes:   []byte("bytes"),
		FStrings: []string{"a", "b"},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response from stream: %v", err)
		}

		// Validate the response
		if resp.FString != req.FString {
			t.Errorf("Expected FString %s, got %s", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Expected FInt32 %d, got %d", req.FInt32, resp.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Expected FEnum %v, got %v", req.FEnum, resp.FEnum)
		}
	}
}
