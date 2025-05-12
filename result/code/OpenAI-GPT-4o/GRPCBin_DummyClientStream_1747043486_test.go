package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to start stream: %v", err)
	}

	// Constructing and sending test messages
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString: "test",
			FInt32:  int32(i),
			FEnum:   grpcbin.DummyMessage_ENUM_1,
		}

		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate the response
	if response.FString != "test" {
		t.Errorf("Expected response FString 'test', got %v", response.FString)
	}

	if response.FInt32 != 9 {
		t.Errorf("Expected response FInt32 9, got %v", response.FInt32)
	}

	if response.FEnum != grpcbin.DummyMessage_ENUM_1 {
		t.Errorf("Expected response FEnum ENUM_1, got %v", response.FEnum)
	}
}
