package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create DummyClientStream: %v", err)
	}

	// Prepare and send 10 DummyMessages, varying their content
	var lastMsg *pb.DummyMessage
	for i := 1; i <= 10; i++ {
		msg := &pb.DummyMessage{
			FString:   "test-string",
			FStrings:  []string{"foo", "bar", "baz"},
			FInt32:    int32(i),
			FInt32S:   []int32{1, 2, 3, int32(i)},
			FEnum:     pb.DummyMessage_ENUM_1,
			FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub:      &pb.DummyMessage_Sub{FString: "sub-string"},
			FSubs:     []*pb.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
			FBool:     i%2 == 0,
			FBools:    []bool{true, false},
			FInt64:    int64(i * 100),
			FInt64S:   []int64{100, 200, int64(i * 100)},
			FBytes:    []byte{0x01, 0x02, 0x03},
			FBytess:   [][]byte{{0x04, 0x05}, {0x06}},
			FFloat:    float32(i) * 1.1,
			FFloats:   []float32{2.1, 3.2, float32(i) * 1.3},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send DummyMessage on stream: %v", err)
		}
		lastMsg = msg
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive final DummyMessage: %v", err)
	}

	// Validate: according to service, reply should echo the last message sent
	if reply == nil {
		t.Fatalf("Expected non-nil reply from server")
	}

	// Validate several fields; for simplicity, only sample a few
	if reply.FString != lastMsg.FString || reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("Reply does not match last sent message: got %+v, want %+v", reply, lastMsg)
	}
	if len(reply.FStrings) != len(lastMsg.FStrings) {
		t.Errorf("Reply FStrings length mismatch: got %d, want %d", len(reply.FStrings), len(lastMsg.FStrings))
	}
	for i, v := range lastMsg.FStrings {
		if reply.FStrings[i] != v {
			t.Errorf("Reply FStrings[%d] mismatch: got %s, want %s", i, reply.FStrings[i], v)
		}
	}
}
