package grpcbin_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestGRPCBin_DummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial grpcb.in:9000: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("DummyClientStream start failed: %v", err)
	}

	// Prepare and send 10 DummyMessages
	var lastMsg *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "Test String",
			FStrings: []string{"A", "B"},
			FInt32:   int32(i),
			FInt32S:  []int32{int32(i), int32(i + 1)},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "suba"}, {FString: "subb"}},
			FBool:    i%2 == 0,
			FBools:   []bool{true, false},
			FInt64:   int64(i * 100),
			FInt64S:  []int64{int64(i), int64(i + 10)},
			FBytes:   []byte{0x1, 0x2, 0x3},
			FBytess:  [][]byte{[]byte("ab"), []byte("cd")},
			FFloat:   float32(i) + 1.5,
			FFloats:  []float32{1.1, 2.2},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("stream.Send(msg %d) failed: %v", i, err)
		}
		lastMsg = msg
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("stream.CloseAndRecv() failed: %v", err)
	}

	// Validate the reply matches the last sent message
	if reply.FString != lastMsg.FString {
		t.Errorf("FString mismatch: got %q, want %q", reply.FString, lastMsg.FString)
	}
	if !equalStringSlice(reply.FStrings, lastMsg.FStrings) {
		t.Errorf("FStrings mismatch: got %v, want %v", reply.FStrings, lastMsg.FStrings)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", reply.FInt32, lastMsg.FInt32)
	}
	if !equalInt32Slice(reply.FInt32S, lastMsg.FInt32S) {
		t.Errorf("FInt32S mismatch: got %v, want %v", reply.FInt32S, lastMsg.FInt32S)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", reply.FEnum, lastMsg.FEnum)
	}
	if !equalEnumSlice(reply.FEnums, lastMsg.FEnums) {
		t.Errorf("FEnums mismatch: got %v, want %v", reply.FEnums, lastMsg.FEnums)
	}
	if (reply.FSub == nil) != (lastMsg.FSub == nil) ||
		(reply.FSub != nil && reply.FSub.FString != lastMsg.FSub.FString) {
		t.Errorf("FSub mismatch: got %+v, want %+v", reply.FSub, lastMsg.FSub)
	}
	if !equalSubSlice(reply.FSubs, lastMsg.FSubs) {
		t.Errorf("FSubs mismatch: got %+v, want %+v", reply.FSubs, lastMsg.FSubs)
	}
	if reply.FBool != lastMsg.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", reply.FBool, lastMsg.FBool)
	}
	if !equalBoolSlice(reply.FBools, lastMsg.FBools) {
		t.Errorf("FBools mismatch: got %v, want %v", reply.FBools, lastMsg.FBools)
	}
	if reply.FInt64 != lastMsg.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", reply.FInt64, lastMsg.FInt64)
	}
	if !equalInt64Slice(reply.FInt64S, lastMsg.FInt64S) {
		t.Errorf("FInt64S mismatch: got %v, want %v", reply.FInt64S, lastMsg.FInt64S)
	}
	if !bytes.Equal(reply.FBytes, lastMsg.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", reply.FBytes, lastMsg.FBytes)
	}
	if !equalBytesSlice(reply.FBytess, lastMsg.FBytess) {
		t.Errorf("FBytess mismatch: got %v, want %v", reply.FBytess, lastMsg.FBytess)
	}
	if reply.FFloat != lastMsg.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", reply.FFloat, lastMsg.FFloat)
	}
	if !equalFloat32Slice(reply.FFloats, lastMsg.FFloats) {
		t.Errorf("FFloats mismatch: got %v, want %v", reply.FFloats, lastMsg.FFloats)
	}
}

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalInt32Slice(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalEnumSlice(a, b []grpcbin.DummyMessage_Enum) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalSubSlice(a, b []*grpcbin.DummyMessage_Sub) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if (a[i] == nil) != (b[i] == nil) {
			return false
		}
		if a[i] != nil && a[i].FString != b[i].FString {
			return false
		}
	}
	return true
}

func equalBoolSlice(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalInt64Slice(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalBytesSlice(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !bytes.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalFloat32Slice(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	const epsilon = 1e-6
	for i := range a {
		if abs(a[i]-b[i]) > epsilon {
			return false
		}
	}
	return true
}

func abs(f float32) float32 {
	if f < 0 {
		return -f
	}
	return f
}
