package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestGRPCBin_DummyClientStream(t *testing.T) {
	// Setting up a connection to the server.
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Create stream for DummyClientStream
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Sending DummyMessages
	for i := 1; i <= 10; i++ {
		message := &grpcbin.DummyMessage{
			FString:  "test",
			FInt32:   int32(i),
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FSub:     &grpcbin.DummyMessage_Sub{FString: "subtest"},
			FBool:    true,
			FInt64:   int64(i),
			FFloat:   1.1,
		}
		if err := stream.Send(message); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	// Receive response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validating the response
	expectedString := "test"
	if resp.FString != expectedString {
		t.Fatalf("Response FString: expected %v, got %v", expectedString, resp.FString)
	}

	expectedInt32 := int32(10)
	if resp.FInt32 != expectedInt32 {
		t.Fatalf("Response FInt32: expected %v, got %v", expectedInt32, resp.FInt32)
	}

	expectedEnum := grpcbin.DummyMessage_ENUM_1
	if resp.FEnum != expectedEnum {
		t.Fatalf("Response FEnum: expected %v, got %v", expectedEnum, resp.FEnum)
	}
}
