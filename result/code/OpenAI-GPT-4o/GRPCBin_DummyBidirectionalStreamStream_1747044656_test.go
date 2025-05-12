package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	addr := "grpcb.in:9000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	incomingMsg := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  42,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
	}

	go func() {
		if err := stream.Send(incomingMsg); err != nil {
			t.Errorf("Failed to send message: %v", err)
		}
		if err := stream.CloseSend(); err != nil {
			t.Errorf("Failed to close send stream: %v", err)
		}
	}()

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	if resp.FString != incomingMsg.FString || resp.FInt32 != incomingMsg.FInt32 || resp.FEnum != incomingMsg.FEnum {
		t.Errorf("Unexpected response: got %v want %v", resp, incomingMsg)
	}
}
