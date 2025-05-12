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
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to finish stream: %v", err)
	}

	messages := []pb.DummyMessage{
		{FString: "test1", FEnum: pb.DummyMessage_ENUM_0},
		{FString: "test2", FEnum: pb.DummyMessage_ENUM_1},
		// ...construct remaining messages
	}

	for _, msg := range messages {
		if err := stream.Send(&msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	if reply.FString != messages[len(messages)-1].FString {
		t.Errorf("Expected %v, got %v", messages[len(messages)-1].FString, reply.FString)
	}
}
