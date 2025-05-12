package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to open stream: %v", err)
	}

	// Construct a typical request
	dummyMessages := []*pb.DummyMessage{
		{
			FString: "test",
			FInt32:  123,
			FEnum:   pb.DummyMessage_ENUM_1,
		},
	}

	for _, req := range dummyMessages {
		if err := stream.Send(req); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}

		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive message: %v", err)
		}

		// Validate client response
		if resp.FString != req.FString || resp.FInt32 != req.FInt32 || resp.FEnum != req.FEnum {
			t.Errorf("Unexpected response: got %v, want %v", resp, req)
		}
	}

	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to close stream: %v", err)
	}
}
