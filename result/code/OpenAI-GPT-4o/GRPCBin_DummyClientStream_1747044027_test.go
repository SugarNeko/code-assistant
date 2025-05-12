package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("could not invoke service: %v", err)
	}

	// Constructing typical requests
	messages := []*grpcbin.DummyMessage{
		{FString: "test1"},
		{FString: "test2"},
		{FString: "test3"},
		// Populate more messages as needed
	}

	for _, msg := range messages {
		if err := stream.Send(msg); err != nil {
			t.Fatalf("could not send message: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("could not receive response: %v", err)
	}

	// Server response validation
	if resp == nil || resp.FString != "test3" {
		t.Errorf("unexpected server response: %v", resp)
	}

	t.Logf("success: received expected response %v", resp)
}
