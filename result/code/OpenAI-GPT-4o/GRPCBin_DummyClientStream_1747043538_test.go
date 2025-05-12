package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Sending 10 DummyMessages
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test",
			FInt32:   int32(i),
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FSub:     &grpcbin.DummyMessage_Sub{FString: "subtest"},
			FBool:    true,
			FInt64:   int64(i),
			FFloat:   float32(i) * 1.1,
			FBytes:   []byte("testbytes"),
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive reply: %v", err)
	}

	// Validate client response
	if reply.FString != "test" {
		t.Errorf("Expected reply.FString to be 'test', got %v", reply.FString)
	}
	if reply.FInt32 != 9 {
		t.Errorf("Expected reply.FInt32 to be 9, got %v", reply.FInt32)
	}
	if reply.FEnum != grpcbin.DummyMessage_ENUM_1 {
		t.Errorf("Expected reply.FEnum to be ENUM_1, got %v", reply.FEnum)
	}
	if reply.FSub.FString != "subtest" {
		t.Errorf("Expected reply.FSub.FString to be 'subtest', got %v", reply.FSub.FString)
	}
	if !reply.FBool {
		t.Errorf("Expected reply.FBool to be true, got %v", reply.FBool)
	}
	if reply.FInt64 != 9 {
		t.Errorf("Expected reply.FInt64 to be 9, got %v", reply.FInt64)
	}
	if reply.FFloat != 9.9 {
		t.Errorf("Expected reply.FFloat to be 9.9, got %v", reply.FFloat)
	}
	if string(reply.FBytes) != "testbytes" {
		t.Errorf("Expected reply.FBytes to be 'testbytes', got %v", string(reply.FBytes))
	}
}
