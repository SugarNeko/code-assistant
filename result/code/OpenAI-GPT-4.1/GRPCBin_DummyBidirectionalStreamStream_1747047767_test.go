package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	// Construct a fully specified DummyMessage
	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"a", "b"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "subfield"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    987654321,
		FInt64S:   []int64{111, 222},
		FBytes:    []byte("bytesval"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    12.34,
		FFloats:   []float32{1.1, 2.2},
	}

	// Send
	if err := stream.Send(req); err != nil {
		t.Fatalf("send failed: %v", err)
	}

	// Receive
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("receive failed: %v", err)
	}

	// Client response validation
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
	}
	for i, v := range req.FStrings {
		if resp.FStrings[i] != v {
			t.Errorf("FStrings[%d] mismatch: got %q, want %q", i, resp.FStrings[i], v)
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %d, want %d", len(resp.FInt32S), len(req.FInt32S))
	}
	for i, v := range req.FInt32S {
		if resp.FInt32S[i] != v {
			t.Errorf("FInt32S[%d] mismatch: got %d, want %d", i, resp.FInt32S[i], v)
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(resp.FEnums), len(req.FEnums))
	}
	for i, v := range req.FEnums {
		if resp.FEnums[i] != v {
			t.Errorf("FEnums[%d] mismatch: got %v, want %v", i, resp.FEnums[i], v)
		}
	}
	if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %q, want %q", resp.FSub.GetFString(), req.FSub.GetFString())
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
	}
	for i, v := range req.FSubs {
		if resp.FSubs[i].FString != v.FString {
			t.Errorf("FSubs[%d].FString mismatch: got %q, want %q", i, resp.FSubs[i].FString, v.FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(resp.FBools), len(req.FBools))
	}
	for i, v := range req.FBools {
		if resp.FBools[i] != v {
			t.Errorf("FBools[%d] mismatch: got %v, want %v", i, resp.FBools[i], v)
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %d, want %d", len(resp.FInt64S), len(req.FInt64S))
	}
	for i, v := range req.FInt64S {
		if resp.FInt64S[i] != v {
			t.Errorf("FInt64S[%d] mismatch: got %v, want %v", i, resp.FInt64S[i], v)
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
	}
	for i, v := range req.FBytess {
		if string(resp.FBytess[i]) != string(v) {
			t.Errorf("FBytess[%d] mismatch: got %s, want %s", i, resp.FBytess[i], v)
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(resp.FFloats), len(req.FFloats))
	}
	for i, v := range req.FFloats {
		if resp.FFloats[i] != v {
			t.Errorf("FFloats[%d] mismatch: got %v, want %v", i, resp.FFloats[i], v)
		}
	}

	// Try to close sending direction
	if err := stream.CloseSend(); err != nil {
		t.Errorf("CloseSend error: %v", err)
	}
}
