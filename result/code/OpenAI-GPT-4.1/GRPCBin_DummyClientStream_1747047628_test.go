package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	var lastMsg *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"a", "b"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "s1"},
				{FString: "s2"},
			},
			FBool:   i%2 == 0,
			FBools:  []bool{true, false},
			FInt64:  int64(i * 10),
			FInt64S: []int64{100, 200},
			FBytes:  []byte("bytes"),
			FBytess: [][]byte{[]byte("b1"), []byte("b2")},
			FFloat:  float32(i) * 1.5,
			FFloats: []float32{2.3, 4.5},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message %d: %v", i, err)
		}
		lastMsg = msg
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive reply: %v", err)
	}

	// Validate the reply matches the last message sent
	if reply.FString != lastMsg.FString ||
		reply.FInt32 != lastMsg.FInt32 ||
		reply.FEnum != lastMsg.FEnum ||
		reply.FBool != lastMsg.FBool ||
		reply.FInt64 != lastMsg.FInt64 ||
		reply.FFloat != lastMsg.FFloat {
		t.Errorf("server reply does not match last sent message\nGot: %+v\nWant: %+v", reply, lastMsg)
	}
}
