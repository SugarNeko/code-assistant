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
		t.Fatalf("failed to open stream: %v", err)
	}

	// Construct and send 10 DummyMessage values
	var lastSent *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "test-string",
			FStrings:  []string{"a", "b", "c"},
			FInt32:    int32(i),
			FInt32S:   []int32{1, 2, 3},
			FEnum:     grpcbin.DummyMessage_ENUM_2,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "nested-string"},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "one"}, {FString: "two"}},
			FBool:     i%2 == 0,
			FBools:    []bool{true, false},
			FInt64:    int64(100 + i),
			FInt64S:   []int64{1001, 1002, 1003},
			FBytes:    []byte("hello"),
			FBytess:   [][]byte{[]byte("a"), []byte("b")},
			FFloat:    3.14,
			FFloats:   []float32{1.2, 3.4, 5.6},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message #%d: %v", i, err)
		}
		lastSent = msg
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	// Validate the server response matches the latest sent message, field by field
	if reply.FString != lastSent.FString {
		t.Errorf("FString mismatch: got %q, want %q", reply.FString, lastSent.FString)
	}
	if len(reply.FStrings) != len(lastSent.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(reply.FStrings), len(lastSent.FStrings))
	}
	if reply.FInt32 != lastSent.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", reply.FInt32, lastSent.FInt32)
	}
	if len(reply.FInt32S) != len(lastSent.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %d, want %d", len(reply.FInt32S), len(lastSent.FInt32S))
	}
	if reply.FEnum != lastSent.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", reply.FEnum, lastSent.FEnum)
	}
	if len(reply.FEnums) != len(lastSent.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(reply.FEnums), len(lastSent.FEnums))
	}
	if reply.FSub == nil || lastSent.FSub == nil || reply.FSub.FString != lastSent.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v, want %v", reply.FSub, lastSent.FSub)
	}
	if len(reply.FSubs) != len(lastSent.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(reply.FSubs), len(lastSent.FSubs))
	}
	if reply.FBool != lastSent.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", reply.FBool, lastSent.FBool)
	}
	if len(reply.FBools) != len(lastSent.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(reply.FBools), len(lastSent.FBools))
	}
	if reply.FInt64 != lastSent.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", reply.FInt64, lastSent.FInt64)
	}
	if len(reply.FInt64S) != len(lastSent.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %d, want %d", len(reply.FInt64S), len(lastSent.FInt64S))
	}
	if string(reply.FBytes) != string(lastSent.FBytes) {
		t.Errorf("FBytes mismatch: got %s, want %s", string(reply.FBytes), string(lastSent.FBytes))
	}
	if len(reply.FBytess) != len(lastSent.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(reply.FBytess), len(lastSent.FBytess))
	}
	if reply.FFloat != lastSent.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", reply.FFloat, lastSent.FFloat)
	}
	if len(reply.FFloats) != len(lastSent.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(reply.FFloats), len(lastSent.FFloats))
	}
}
