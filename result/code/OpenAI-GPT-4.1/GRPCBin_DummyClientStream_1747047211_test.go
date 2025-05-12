package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	// Connect to grpcb.in:9000 with 15s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	// Construct 10 DummyMessages per proto specification
	var lastMsg *pb.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString:   "test-string",
			FStrings:  []string{"foo", "bar"},
			FInt32:    int32(i),
			FInt32S:   []int32{int32(i), int32(i + 1)},
			FEnum:     pb.DummyMessage_ENUM_1,
			FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub:      &pb.DummyMessage_Sub{FString: "sub-string"},
			FSubs:     []*pb.DummyMessage_Sub{{FString: "rep-sub-1"}, {FString: "rep-sub-2"}},
			FBool:     true,
			FBools:    []bool{true, false},
			FInt64:    int64(i * 100),
			FInt64S:   []int64{int64(i * 10), int64(i * 20)},
			FBytes:    []byte("bytes"),
			FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
			FFloat:    1.23,
			FFloats:   []float32{4.56, 7.89},
		}
		lastMsg = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message %d: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive reply: %v", err)
	}

	// Validate that the server echoed back the last sent message
	if reply == nil {
		t.Fatalf("expected reply, got nil")
	}
	if reply.FString != lastMsg.FString ||
		reply.FInt32 != lastMsg.FInt32 ||
		reply.FEnum != lastMsg.FEnum ||
		reply.FSub.FString != lastMsg.FSub.FString ||
		reply.FBool != lastMsg.FBool ||
		reply.FInt64 != lastMsg.FInt64 ||
		string(reply.FBytes) != string(lastMsg.FBytes) ||
		reply.FFloat != lastMsg.FFloat {
		t.Errorf("reply does not match last sent message: got %+v, want %+v", reply, lastMsg)
	}
}
