package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("could not start stream: %v", err)
	}

	// Send test messages
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString: "test",
			FInt32:  int32(i),
			FEnum:   grpcbin.DummyMessage_ENUM_1,
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("could not send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("could not receive reply: %v", err)
	}

	// Validate the response
	expected := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  int32(9),
		FEnum:   grpcbin.DummyMessage_ENUM_1,
	}

	if reply.GetFString() != expected.GetFString() || reply.GetFInt32() != expected.GetFInt32() || reply.GetFEnum() != expected.GetFEnum() {
		t.Errorf("unexpected reply: got %v, want %v", reply, expected)
	}
}
