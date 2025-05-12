package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin" // Adjust the import path if necessary
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("could not create stream: %v", err)
	}

	// Send a dummy message
	req := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  1,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("could not send message: %v", err)
	}

	// Receive a response
	res, err := stream.Recv()
	if err != nil {
		t.Fatalf("could not receive message: %v", err)
	}

	// Validate the response
	if res.FString != req.FString {
		t.Errorf("expected FString %q, got %q", req.FString, res.FString)
	}
	if res.FInt32 != req.FInt32 {
		t.Errorf("expected FInt32 %d, got %d", req.FInt32, res.FInt32)
	}
	if res.FEnum != req.FEnum {
		t.Errorf("expected FEnum %v, got %v", req.FEnum, res.FEnum)
	}
}
