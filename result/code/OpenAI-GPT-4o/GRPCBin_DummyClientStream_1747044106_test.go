package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
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

	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "test",
			FInt32:    int32(i),
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
			FInt64:    int64(i * 100),
			FBytes:    []byte("byteTest"),
			FFloat:    float32(i) * 1.5,
			FBool:     i%2 == 0,
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive reply: %v", err)
	}

	if reply.FString != "test" || reply.FInt32 != 9 || reply.FEnum != grpcbin.DummyMessage_ENUM_1 || reply.FSub.FString != "sub" || reply.FInt64 != 900 || string(reply.FBytes) != "byteTest" || reply.FFloat != 13.5 {
		t.Errorf("Unexpected reply: %+v", reply)
	}
}
