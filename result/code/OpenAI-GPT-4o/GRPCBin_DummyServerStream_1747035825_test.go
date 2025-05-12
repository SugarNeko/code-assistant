package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	// Define the server address and timeout
	const (
		address = "grpcb.in:9000"
		timeout = 15 * time.Second
	)

	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(timeout))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Set up context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Define a valid request
	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   42,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "subTest"},
		FInt64:   123456789,
		FBool:    true,
		FFloat:   1.23,
		FStrings: []string{"one", "two"},
	}

	// Call the DummyServerStream method
	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream call failed: %v", err)
	}

	// Validate responses
	expectedResponses := 10
	for i := 0; i < expectedResponses; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("failed to receive response: %v", err)
		}

		// Check response data
		if resp.FString != req.FString || resp.FInt32 != req.FInt32 {
			t.Errorf("unexpected response: got %v, want %v", resp, req)
		}
	}

	// Verify if the stream ends properly
	_, err = stream.Recv()
	if err == nil {
		t.Errorf("expected stream to end but received additional response")
	}
}
