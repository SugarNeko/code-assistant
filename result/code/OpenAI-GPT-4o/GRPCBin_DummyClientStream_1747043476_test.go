package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString:  "Test String",
			FInt32:   int32(i),
			FEnum:    pb.DummyMessage_ENUM_1,
			FSub:     &pb.DummyMessage_Sub{FString: "Sub Test"},
			FBool:    true,
			FInt64:   int64(i),
			FBytes:   []byte("Test Bytes"),
			FFloat:   float32(i),
		}

		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive reply: %v", err)
	}

	if reply.FString != "Test String" || reply.FInt32 != 9 || reply.FBool != true {
		t.Errorf("Unexpected response received: %v", reply)
	}
}
