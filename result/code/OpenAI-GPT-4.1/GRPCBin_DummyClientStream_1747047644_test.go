package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
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
		t.Fatalf("DummyClientStream failed: %v", err)
	}

	var lastMessage *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:    "msg-" + string(rune('A'+i)),
			FStrings:   []string{"a", "b", "c"},
			FInt32:     int32(i),
			FInt32S:    []int32{int32(i), int32(i + 1)},
			FEnum:      grpcbin.DummyMessage_ENUM_1,
			FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub:       &grpcbin.DummyMessage_Sub{FString: "sub-string"},
			FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:      i%2 == 0,
			FBools:     []bool{true, false},
			FInt64:     int64(i * 10),
			FInt64S:    []int64{int64(i), int64(i + 5)},
			FBytes:     []byte{0xAB, 0xCD},
			FBytess:    [][]byte{[]byte("foo"), []byte("bar")},
			FFloat:     1.23 + float32(i),
			FFloats:    []float32{2.34, 3.45},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
		lastMessage = msg
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("CloseAndRecv failed: %v", err)
	}

	// Validate the response matches the last sent message
	if reply.FString != lastMessage.FString {
		t.Errorf("reply.FString = %q; want %q", reply.FString, lastMessage.FString)
	}
	if reply.FInt32 != lastMessage.FInt32 {
		t.Errorf("reply.FInt32 = %d; want %d", reply.FInt32, lastMessage.FInt32)
	}
	if reply.FEnum != lastMessage.FEnum {
		t.Errorf("reply.FEnum = %v; want %v", reply.FEnum, lastMessage.FEnum)
	}
	if reply.FBool != lastMessage.FBool {
		t.Errorf("reply.FBool = %v; want %v", reply.FBool, lastMessage.FBool)
	}
	if reply.FInt64 != lastMessage.FInt64 {
		t.Errorf("reply.FInt64 = %d; want %d", reply.FInt64, lastMessage.FInt64)
	}
	if len(reply.FStrings) != len(lastMessage.FStrings) {
		t.Errorf("reply.FStrings len = %d; want %d", len(reply.FStrings), len(lastMessage.FStrings))
	}
}
