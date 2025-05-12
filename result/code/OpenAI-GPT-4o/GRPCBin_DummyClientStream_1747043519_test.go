package grpcbin_test

import (
	"context"
	"testing"
	"time"
	
	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Client stream failed: %v", err)
	}

	testMessages := []*grpcbin.DummyMessage{
		{FString: "test1"},
		{FString: "test2"},
	}

	for _, msg := range testMessages {
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive message: %v", err)
	}

	expectedFString := "test2"
	if reply.FString != expectedFString {
		t.Errorf("expected %q, got %q", expectedFString, reply.FString)
	}
}
