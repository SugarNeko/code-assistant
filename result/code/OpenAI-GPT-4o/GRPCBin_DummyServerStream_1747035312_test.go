package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct request
	req := &grpcbin.DummyMessage{
		FString: "test_string",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("could not get stream: %v", err)
	}

	// Read server responses
	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("failed to receive response: %v", err)
		}

		// Validate the response
		if resp.FString != req.FString {
			t.Errorf("unexpected response FString: got %v, want %v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32*10 {
			t.Errorf("unexpected response FInt32: got %v, want %v", resp.FInt32, req.FInt32*10)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("unexpected response FEnum: got %v, want %v", resp.FEnum, req.FEnum)
		}
	}

	log.Println("TestDummyServerStream passed")
}
