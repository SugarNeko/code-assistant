package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	request := &pb.DummyMessage{
		FString:  "Test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FBoolean: true,
	}

	if err := stream.Send(request); err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	res, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	expected := *request
	if *res != expected {
		t.Errorf("Response did not match expected value. Got %v, want %v", res, expected)
	}
}
