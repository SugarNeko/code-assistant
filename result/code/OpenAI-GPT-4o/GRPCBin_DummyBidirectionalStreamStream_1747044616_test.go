package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	message := &grpcbin.DummyMessage{
		FString: "test",
		FEnum:   grpcbin.DummyMessage_ENUM_1,
	}

	if err := stream.Send(message); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	if resp.GetFString() != message.GetFString() {
		t.Errorf("Expected %v, got %v", message.GetFString(), resp.GetFString())
	}

	if resp.GetFEnum() != message.GetFEnum() {
		t.Errorf("Expected %v, got %v", message.GetFEnum(), resp.GetFEnum())
	}
}
