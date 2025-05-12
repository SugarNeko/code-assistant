package grpcbin_test

import (
	"context"
	"testing"
	"time"
	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("could not open stream: %v", err)
	}

	// Send 10 DummyMessages
	for i := 0; i < 10; i++ {
		message := &pb.DummyMessage{
			FString: "Test Message",
		}
		if err := stream.Send(message); err != nil {
			t.Fatalf("could not send message %v: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("could not receive reply: %v", err)
	}

	// Validate response
	if reply.GetFString() != "Test Message" {
		t.Errorf("unexpected response: got %v, want %v", reply.GetFString(), "Test Message")
	}
}
