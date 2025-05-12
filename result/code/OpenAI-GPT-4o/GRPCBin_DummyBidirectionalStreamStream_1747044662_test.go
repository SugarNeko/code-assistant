package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
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
		FString: "test string",
		FInt32:  42,
		FEnum:   pb.DummyMessage_ENUM_1,
		FSub:    &pb.DummyMessage_Sub{FString: "sub string"},
		FBool:   true,
		FInt64:  64,
		FFloat:  3.14,
	}

	if err := stream.Send(request); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	response, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	if response.FString != request.FString {
		t.Errorf("Expected FString to be %q, got %q", request.FString, response.FString)
	}

	if response.FInt32 != request.FInt32 {
		t.Errorf("Expected FInt32 to be %d, got %d", request.FInt32, response.FInt32)
	}

	if response.FEnum != request.FEnum {
		t.Errorf("Expected FEnum to be %v, got %v", request.FEnum, response.FEnum)
	}

	if response.FBool != request.FBool {
		t.Errorf("Expected FBool to be %t, got %t", request.FBool, response.FBool)
	}

	if response.FInt64 != request.FInt64 {
		t.Errorf("Expected FInt64 to be %d, got %d", request.FInt64, response.FInt64)
	}

	if response.FFloat != request.FFloat {
		t.Errorf("Expected FFloat to be %f, got %f", request.FFloat, response.FFloat)
	}
}
