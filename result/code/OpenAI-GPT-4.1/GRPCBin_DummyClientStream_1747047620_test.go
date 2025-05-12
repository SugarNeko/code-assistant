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
		t.Fatalf("DummyClientStream failed to start: %v", err)
	}

	// Construct and send 10 DummyMessage
	var lastMsg *grpcbin.DummyMessage
	for i := 1; i <= 10; i++ {
		m := &grpcbin.DummyMessage{
			FString:   "message_" + string(rune(i+'0')),
			FStrings:  []string{"foo", "bar"},
			FInt32:    int32(i),
			FInt32S:   []int32{1, 2, 3},
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-" + string(rune(i+'0'))},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     i%2 == 0,                            // alternate true/false
			FBools:    []bool{true, false, true},
			FInt64:    int64(i * 100),
			FInt64S:   []int64{100, 200, 300},
			FBytes:    []byte{'a', 'b', 'c'},
			FBytess:   [][]byte{{'x', 'y'}, {'a', 'b'}},
			FFloat:    float32(i) + 0.5,
			FFloats:   []float32{1.1, 2.2},
		}
		lastMsg = m
		if err := stream.Send(m); err != nil {
			t.Fatalf("failed to send DummyMessage %d: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive DummyMessage: %v", err)
	}

	// Validate server response (should match the last message sent)
	if reply.FString != lastMsg.FString {
		t.Errorf("expected FString=%s, got %s", lastMsg.FString, reply.FString)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("expected FInt32=%d, got %d", lastMsg.FInt32, reply.FInt32)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("expected FEnum=%v, got %v", lastMsg.FEnum, reply.FEnum)
	}
	if reply.FBool != lastMsg.FBool {
		t.Errorf("expected FBool=%v, got %v", lastMsg.FBool, reply.FBool)
	}
	if reply.FInt64 != lastMsg.FInt64 {
		t.Errorf("expected FInt64=%d, got %d", lastMsg.FInt64, reply.FInt64)
	}
	if string(reply.FBytes) != string(lastMsg.FBytes) {
		t.Errorf("expected FBytes=%v, got %v", lastMsg.FBytes, reply.FBytes)
	}
	if reply.FFloat != lastMsg.FFloat {
		t.Errorf("expected FFloat=%f, got %f", lastMsg.FFloat, reply.FFloat)
	}
}
