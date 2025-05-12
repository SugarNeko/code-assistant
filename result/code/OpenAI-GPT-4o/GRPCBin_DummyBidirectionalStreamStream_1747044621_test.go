package grpcbin_test

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
		t.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Error creating stream: %v", err)
	}

	sentMessage := &grpcbin.DummyMessage{
		FString: "Test",
		FInt32:  42,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
	}

	err = stream.Send(sentMessage)
	if err != nil {
		t.Fatalf("Error sending message: %v", err)
	}

	receivedMessage, err := stream.Recv()
	if err != nil {
		t.Fatalf("Error receiving message: %v", err)
	}

	if receivedMessage.FString != sentMessage.FString {
		t.Errorf("Expected FString: %v, got: %v", sentMessage.FString, receivedMessage.FString)
	}

	if receivedMessage.FInt32 != sentMessage.FInt32 {
		t.Errorf("Expected FInt32: %v, got: %v", sentMessage.FInt32, receivedMessage.FInt32)
	}

	if receivedMessage.FEnum != sentMessage.FEnum {
		t.Errorf("Expected FEnum: %v, got: %v", sentMessage.FEnum, receivedMessage.FEnum)
	}
}
