package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	// Set timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Dial to grpcb.in:9000 with timeout
	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to start stream: %v", err)
	}

	// Construct a full DummyMessage request as per proto spec
	in := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_2,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "a"}, {FString: "b"}},
		FBool:    true,
		FBools:   []bool{false, true},
		FInt64:   99,
		FInt64S:  []int64{7, 8},
		FBytes:   []byte("hello"),
		FBytess:  [][]byte{[]byte("a"), []byte("b")},
		FFloat:   1.23,
		FFloats:  []float32{4.56, 7.89},
	}

	// Send message to stream
	if err := stream.Send(in); err != nil {
		t.Fatalf("Failed to send DummyMessage: %v", err)
	}
	// Receive echo message from stream
	out, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive DummyMessage: %v", err)
	}

	// Validate the output equals input
	if out.FString != in.FString {
		t.Errorf("FString: got %q, want %q", out.FString, in.FString)
	}
	if len(out.FStrings) != len(in.FStrings) {
		t.Errorf("FStrings: length mismatch: got %d, want %d", len(out.FStrings), len(in.FStrings))
	}
	for i := range in.FStrings {
		if out.FStrings[i] != in.FStrings[i] {
			t.Errorf("FStrings[%d]: got %q, want %q", i, out.FStrings[i], in.FStrings[i])
		}
	}
	if out.FInt32 != in.FInt32 {
		t.Errorf("FInt32: got %d, want %d", out.FInt32, in.FInt32)
	}
	if len(out.FInt32S) != len(in.FInt32S) {
		t.Errorf("FInt32S: length mismatch")
	}
	for i := range in.FInt32S {
		if out.FInt32S[i] != in.FInt32S[i] {
			t.Errorf("FInt32S[%d]: got %d, want %d", i, out.FInt32S[i], in.FInt32S[i])
		}
	}
	if out.FEnum != in.FEnum {
		t.Errorf("FEnum: got %v, want %v", out.FEnum, in.FEnum)
	}
	if len(out.FEnums) != len(in.FEnums) {
		t.Errorf("FEnums: length mismatch")
	}
	for i := range in.FEnums {
		if out.FEnums[i] != in.FEnums[i] {
			t.Errorf("FEnums[%d]: got %v, want %v", i, out.FEnums[i], in.FEnums[i])
		}
	}
	if out.FSub == nil || in.FSub == nil {
		t.Errorf("FSub: nil value(s)")
	} else if out.FSub.FString != in.FSub.FString {
		t.Errorf("FSub.FString: got %q, want %q", out.FSub.FString, in.FSub.FString)
	}
	if len(out.FSubs) != len(in.FSubs) {
		t.Errorf("FSubs: length mismatch")
	}
	for i := range in.FSubs {
		if out.FSubs[i].FString != in.FSubs[i].FString {
			t.Errorf("FSubs[%d].FString: got %q, want %q", i, out.FSubs[i].FString, in.FSubs[i].FString)
		}
	}
	if out.FBool != in.FBool {
		t.Errorf("FBool: got %v, want %v", out.FBool, in.FBool)
	}
	if len(out.FBools) != len(in.FBools) {
		t.Errorf("FBools: length mismatch")
	}
	for i := range in.FBools {
		if out.FBools[i] != in.FBools[i] {
			t.Errorf("FBools[%d]: got %v, want %v", i, out.FBools[i], in.FBools[i])
		}
	}
	if out.FInt64 != in.FInt64 {
		t.Errorf("FInt64: got %d, want %d", out.FInt64, in.FInt64)
	}
	if len(out.FInt64S) != len(in.FInt64S) {
		t.Errorf("FInt64S: length mismatch")
	}
	for i := range in.FInt64S {
		if out.FInt64S[i] != in.FInt64S[i] {
			t.Errorf("FInt64S[%d]: got %d, want %d", i, out.FInt64S[i], in.FInt64S[i])
		}
	}
	if string(out.FBytes) != string(in.FBytes) {
		t.Errorf("FBytes: got %v, want %v", out.FBytes, in.FBytes)
	}
	if len(out.FBytess) != len(in.FBytess) {
		t.Errorf("FBytess: length mismatch")
	}
	for i := range in.FBytess {
		if string(out.FBytess[i]) != string(in.FBytess[i]) {
			t.Errorf("FBytess[%d]: got %v, want %v", i, out.FBytess[i], in.FBytess[i])
		}
	}
	if out.FFloat != in.FFloat {
		t.Errorf("FFloat: got %v, want %v", out.FFloat, in.FFloat)
	}
	if len(out.FFloats) != len(in.FFloats) {
		t.Errorf("FFloats: length mismatch")
	}
	for i := range in.FFloats {
		if out.FFloats[i] != in.FFloats[i] {
			t.Errorf("FFloats[%d]: got %v, want %v", i, out.FFloats[i], in.FFloats[i])
		}
	}

	// Close the send direction and end test
	if err := stream.CloseSend(); err != nil {
		t.Errorf("CloseSend error: %v", err)
	}
}
