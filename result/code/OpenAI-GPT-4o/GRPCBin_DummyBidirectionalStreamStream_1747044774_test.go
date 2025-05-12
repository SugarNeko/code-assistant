package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Error initiating stream: %v", err)
	}

	req := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  1,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "sub"},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Error sending request: %v", err)
	}

	res, err := stream.Recv()
	if err != nil {
		t.Fatalf("Error receiving response: %v", err)
	}

	if res.FString != req.FString ||
		res.FInt32 != req.FInt32 ||
		res.FEnum != req.FEnum {
		t.Errorf("Response validation failed. Expected %#v, got %#v", req, res)
	}

	if err := stream.CloseSend(); err != nil {
		t.Errorf("Error closing stream: %v", err)
	}
}
