package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	address := "grpcb.in:9000"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  123,
		FBool:   true,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	for i := 0; i < 10; i++ {
		res, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		// Validate response
		expectedString := req.FString
		expectedInt32 := req.FInt32
		expectedBool := req.FBool

		if res.FString != expectedString || res.FInt32 != expectedInt32 || res.FBool != expectedBool {
			t.Errorf("Unexpected response: got %v, want %v", res, req)
		}
	}
}
