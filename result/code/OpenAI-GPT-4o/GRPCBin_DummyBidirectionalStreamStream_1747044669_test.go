package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	// Set up a connection to the server with a timeout
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Send a DummyMessage
	req := &pb.DummyMessage{
		FString: "test",
		FInt32:  123,
		FEnum:   pb.DummyMessage_ENUM_1,
		FSub:    &pb.DummyMessage_Sub{FString: "subtest"},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	// Receive and validate the response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	if resp.FString != req.FString || resp.FInt32 != req.FInt32 || resp.FEnum != req.FEnum {
		t.Errorf("Response does not match request, got %v, want %v", resp, req)
	}

	if resp.FSub.GetFString() != req.FSub.GetFString() {
		t.Errorf("Response Sub does not match request, got %v, want %v", resp.FSub.GetFString(), req.FSub.GetFString())
	}

	// Further response validation can be performed as needed
	stream.CloseSend()
}
