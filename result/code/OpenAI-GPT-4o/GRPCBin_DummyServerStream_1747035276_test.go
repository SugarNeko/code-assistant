package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test",
		// Populating other fields...
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("could not get stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("failed to receive: %v", err)
		}

		// Validate response (example validation)
		if resp.FString != req.FString {
			t.Errorf("expected %v, got %v", req.FString, resp.FString)
		}
	}

	_, err = stream.Recv()
	if err == nil {
		t.Fatalf("expected EOF, got additional response")
	}
}
