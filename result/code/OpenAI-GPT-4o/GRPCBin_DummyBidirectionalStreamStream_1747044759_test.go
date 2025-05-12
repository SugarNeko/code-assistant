package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := grpcbin.NewGRPCBinClient(conn)

	// Create a stream
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("could not open stream: %v", err)
	}

	// Prepare and send a request
	req := &grpcbin.DummyMessage{
		FString: "hello",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
	}
	if err := stream.Send(req); err != nil {
		t.Fatalf("could not send request: %v", err)
	}

	// Receive and validate response
	res, err := stream.Recv()
	if err != nil {
		t.Fatalf("could not receive response: %v", err)
	}

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
