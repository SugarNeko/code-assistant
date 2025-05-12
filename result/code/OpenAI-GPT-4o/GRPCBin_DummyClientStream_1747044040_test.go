package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Send 10 DummyMessages
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test",
			FInt32:   int32(i),
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FSub:     &grpcbin.DummyMessage_Sub{FString: "subtest"},
			FBool:    true,
			FInt64:   int64(i),
			FBytes:   []byte("byteTest"),
			FFloat:   float32(i),
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	// Close the stream and receive the last message
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate the response message
	if reply.FString != "test" || reply.FInt32 != 9 {
		t.Fatalf("Unexpected response data: %v", reply)
	}
}
