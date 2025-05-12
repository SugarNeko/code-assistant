package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	address := "grpcb.in:9000"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Send a message to the server
	message := &grpcbin.DummyMessage{
		FString: "test_message",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FBool:   true,
	}
	if err := stream.Send(message); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive a message from the server
	response, err := stream.Recv()
	if err != nil {
		t.Fatalf("Error while receiving response: %v", err)
	}

	// Validate the response
	expectedMessage := message.FString
	if response.FString != expectedMessage {
		t.Errorf("Expected response %v, but got %v", expectedMessage, response.FString)
	}
}
