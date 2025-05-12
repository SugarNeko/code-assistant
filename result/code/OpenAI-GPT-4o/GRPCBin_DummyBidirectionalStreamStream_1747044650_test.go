package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
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

	req := &pb.DummyMessage{
		FString: "test",
		FInt32: 123,
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	expectedResp := &pb.DummyMessage{
		FString: "test",
		FInt32: 123,
	}

	if resp.FString != expectedResp.FString || resp.FInt32 != expectedResp.FInt32 {
		t.Fatalf("Received incorrect response: got %v want %v", resp, expectedResp)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Failed to close stream: %v", err)
	}
}
