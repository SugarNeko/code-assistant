package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Send a valid DummyMessage
	req := &grpcbin.DummyMessage{
		FString: "hello",
		FInt32:  42,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "sub message"},
		FBool:   true,
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive response
	res, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Validate the response
	if res.FString != req.FString {
		t.Errorf("Expected FString: %s, got: %s", req.FString, res.FString)
	}
	if res.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32: %d, got: %d", req.FInt32, res.FInt32)
	}
	if res.FEnum != req.FEnum {
		t.Errorf("Expected FEnum: %v, got: %v", req.FEnum, res.FEnum)
	}
	if res.FSub.FString != req.FSub.FString {
		t.Errorf("Expected FSub FString: %s, got: %s", req.FSub.FString, res.FSub.FString)
	}
	if res.FBool != req.FBool {
		t.Errorf("Expected FBool: %v, got: %v", req.FBool, res.FBool)
	}
}
