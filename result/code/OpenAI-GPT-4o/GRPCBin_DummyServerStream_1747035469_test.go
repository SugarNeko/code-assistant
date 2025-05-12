package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
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
		FInt32:   42,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub"},
		FBool:    true,
		FInt64:   64,
		FFloat:   3.14,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	expectedResponses := 10
	for i := 0; i < expectedResponses; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}
		if resp.FString != req.FString || resp.FInt32 != req.FInt32 {
			t.Errorf("Unexpected response: got %v, expected %v", resp, req)
		}
	}
}
