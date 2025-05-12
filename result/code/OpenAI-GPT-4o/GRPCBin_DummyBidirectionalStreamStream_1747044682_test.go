package grpcbintest

import (
	"context"
	"log"
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
		t.Fatalf("Error creating stream: %v", err)
	}

	req := &pb.DummyMessage{
		FString:   "Test",
		FInt32:    123,
		FEnum:     pb.DummyMessage_ENUM_1,
		FSub:      &pb.DummyMessage_Sub{FString: "SubTest"},
		FBool:     true,
		FInt64:    999,
		FFloat:    1.23,
		FBytes:    []byte("TestBytes"),
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send a message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive a message: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString: %v, got: %v", req.FString, resp.FString)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32: %v, got: %v", req.FInt32, resp.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("Expected FEnum: %v, got: %v", req.FEnum, resp.FEnum)
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("Expected FSub FString: %v, got: %v", req.FSub.FString, resp.FSub.FString)
	}
	if resp.FBool != req.FBool {
		t.Errorf("Expected FBool: %v, got: %v", req.FBool, resp.FBool)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("Expected FInt64: %v, got: %v", req.FInt64, resp.FInt64)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("Expected FBytes: %v, got: %v", req.FBytes, resp.FBytes)
	}
}
