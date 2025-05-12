package grpcbintest

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Send a test message
	testMessage := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   42,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FBool:    true,
		FInt64:   1234567890,
		FFloat:   3.14,
	}

	if err := stream.Send(testMessage); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive the message back
	response, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Validate the received message
	if response.FString != testMessage.FString ||
		response.FInt32 != testMessage.FInt32 ||
		response.FEnum != testMessage.FEnum ||
		response.FSub.FString != testMessage.FSub.FString ||
		response.FBool != testMessage.FBool ||
		response.FInt64 != testMessage.FInt64 ||
		response.FFloat != testMessage.FFloat {
		t.Errorf("Response does not match sent message: got %+v, want %+v", response, testMessage)
	}
}
