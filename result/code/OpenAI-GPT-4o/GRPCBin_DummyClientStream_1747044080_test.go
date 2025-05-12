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
		t.Fatalf("Failed to open stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString: "test",
			FInt32:  int32(i),
			FEnum:   grpcbin.DummyMessage_ENUM_0,
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive reply: %v", err)
	}

	if reply.FString != "test" || reply.FInt32 != 9 {
		t.Errorf("Unexpected reply: %v", reply)
	}
}
