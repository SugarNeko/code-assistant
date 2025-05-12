package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
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
		t.Fatalf("Failed to create stream: %v", err)
	}

	inputMessages := []pb.DummyMessage{
		{FString: "test1", FInt32: 1},
		{FString: "test2", FInt32: 2},
		{FString: "test3", FInt32: 3},
		{FString: "test4", FInt32: 4},
		{FString: "test5", FInt32: 5},
		{FString: "test6", FInt32: 6},
		{FString: "test7", FInt32: 7},
		{FString: "test8", FInt32: 8},
		{FString: "test9", FInt32: 9},
		{FString: "test10", FInt32: 10},
	}

	for _, msg := range inputMessages {
		if err := stream.Send(&msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	receivedMessage, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive final message: %v", err)
	}

	expectedMessage := pb.DummyMessage{FString: "test10", FInt32: 10}
	if receivedMessage.FString != expectedMessage.FString || receivedMessage.FInt32 != expectedMessage.FInt32 {
		t.Errorf("Unexpected response. Got %v, want %v", receivedMessage, expectedMessage)
	}
}
