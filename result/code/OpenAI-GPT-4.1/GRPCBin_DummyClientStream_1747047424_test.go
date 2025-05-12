package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("failed to open client stream: %v", err)
	}

	var lastReq *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "test" + string('A'+i),
			FStrings:  []string{"foo", "bar", "baz"},
			FInt32:    int32(i),
			FInt32S:   []int32{int32(i), int32(i * 2)},
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub" + string('A'+i)},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     i%2 == 0,
			FBools:    []bool{true, false, true},
			FInt64:    int64(i * 100),
			FInt64S:   []int64{10, 20, 30},
			FBytes:    []byte("mybytes"),
			FBytess:   [][]byte{[]byte("a"), []byte("b")},
			FFloat:    float32(i) * 1.5,
			FFloats:   []float32{1.1, 2.2},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message #%d: %v", i, err)
		}
		lastReq = msg
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	// Validate that response matches the last message sent
	if reply.FString != lastReq.FString {
		t.Errorf("response FString mismatch: got %s, want %s", reply.FString, lastReq.FString)
	}
	if reply.FInt32 != lastReq.FInt32 {
		t.Errorf("response FInt32 mismatch: got %d, want %d", reply.FInt32, lastReq.FInt32)
	}
	if len(reply.FStrings) != len(lastReq.FStrings) {
		t.Errorf("response FStrings length mismatch: got %d, want %d", len(reply.FStrings), len(lastReq.FStrings))
	}
	if reply.FEnum != lastReq.FEnum {
		t.Errorf("response FEnum mismatch: got %v, want %v", reply.FEnum, lastReq.FEnum)
	}
	if reply.FSub == nil || reply.FSub.FString != lastReq.FSub.FString {
		t.Errorf("response Sub FString mismatch: got %v, want %v", reply.FSub, lastReq.FSub)
	}
	if reply.FBool != lastReq.FBool {
		t.Errorf("response FBool mismatch: got %t, want %t", reply.FBool, lastReq.FBool)
	}
	if reply.FInt64 != lastReq.FInt64 {
		t.Errorf("response FInt64 mismatch: got %d, want %d", reply.FInt64, lastReq.FInt64)
	}
	if string(reply.FBytes) != string(lastReq.FBytes) {
		t.Errorf("response FBytes mismatch: got %s, want %s", reply.FBytes, lastReq.FBytes)
	}
	if reply.FFloat != lastReq.FFloat {
		t.Errorf("response FFloat mismatch: got %f, want %f", reply.FFloat, lastReq.FFloat)
	}
}
