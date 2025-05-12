package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Error when calling DummyClientStream: %v", err)
	}

	// Send 10 DummyMessages
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString: "test",
			FInt32:  int32(i),
			FEnum:   pb.DummyMessage_ENUM_1,
			FSub:    &pb.DummyMessage_Sub{FString: "sub"},
			FBool:   true,
			FInt64:  int64(i),
			FBool:   true,
			FBytes:  []byte("bytes"),
			FFloat:  1.23,
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Error sending message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Error receiving reply: %v", err)
	}

	if reply.FString != "test" || reply.FInt32 != 9 {
		t.Errorf("Unexpected reply: %v", reply)
	}

	log.Printf("Received: %v", reply)
}
