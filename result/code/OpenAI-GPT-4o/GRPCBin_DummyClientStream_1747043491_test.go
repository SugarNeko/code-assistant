package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	address := "grpcb.in:9000"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	messages := []*pb.DummyMessage{
		{FString: "Test1"},
		{FString: "Test2"},
		{FString: "Test3"},
		{FString: "Test4"},
		{FString: "Test5"},
		{FString: "Test6"},
		{FString: "Test7"},
		{FString: "Test8"},
		{FString: "Test9"},
		{FString: "Test10"},
	}

	for _, msg := range messages {
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive reply: %v", err)
	}

	expected := "Test10"
	if reply.FString != expected {
		t.Errorf("unexpected server response: got %v, want %v", reply.FString, expected)
	}
}
