package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Constructing the request
	req := &pb.DummyMessage{
		FString: "test",
		FInt32:  32,
		FEnum:   pb.DummyMessage_ENUM_1,
		FSub:    &pb.DummyMessage_Sub{FString: "sub"},
		FBool:   true,
		FInt64:  64,
		FFloat:  1.23,
	}

	// Make the gRPC call
	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("could not call DummyServerStream: %v", err)
	}

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("error receiving response: %v", err)
		}

		// Validate the response
		if resp.FString != "test" {
			t.Errorf("Expected FString 'test', got %s", resp.FString)
		}
		if resp.FInt32 != 320 {
			t.Errorf("Expected FInt32 320, got %d", resp.FInt32)
		}
		if resp.FEnum != pb.DummyMessage_ENUM_1 {
			t.Errorf("Expected FEnum ENUM_1, got %v", resp.FEnum)
		}
		if resp.FSub.FString != "sub" {
			t.Errorf("Expected FSub FString 'sub', got %s", resp.FSub.FString)
		}
		if resp.FBool != true {
			t.Errorf("Expected FBool true, got %v", resp.FBool)
		}
		if resp.FInt64 != 640 {
			t.Errorf("Expected FInt64 640, got %d", resp.FInt64)
		}
		if resp.FFloat != 12.3 {
			t.Errorf("Expected FFloat 12.3, got %f", resp.FFloat)
		}
	}
}
