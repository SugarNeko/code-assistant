package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	// Set timeout for connection
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to start bidirectional stream: %v", err)
	}

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub_str"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    1234567890,
		FInt64S:   []int64{100, 200, 300},
		FBytes:    []byte("byte_string"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	// Send request
	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive echo back from server
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Validate response matches request
	if resp.FString != req.FString {
		t.Errorf("FString: got %q, want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings: got length %d, want %d", len(resp.FStrings), len(req.FStrings))
	}
	for i := range req.FStrings {
		if resp.FStrings[i] != req.FStrings[i] {
			t.Errorf("FStrings[%d]: got %q, want %q", i, resp.FStrings[i], req.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32: got %d, want %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S: got length %d, want %d", len(resp.FInt32S), len(req.FInt32S))
	}
	for i := range req.FInt32S {
		if resp.FInt32S[i] != req.FInt32S[i] {
			t.Errorf("FInt32S[%d]: got %d, want %d", i, resp.FInt32S[i], req.FInt32S[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums: got length %d, want %d", len(resp.FEnums), len(req.FEnums))
	}
	for i := range req.FEnums {
		if resp.FEnums[i] != req.FEnums[i] {
			t.Errorf("FEnums[%d]: got %v, want %v", i, resp.FEnums[i], req.FEnums[i])
		}
	}
	if req.FSub != nil && (resp.FSub == nil || resp.FSub.FString != req.FSub.FString) {
		t.Errorf("FSub.FString: got %q, want %q", resp.FSub.GetFString(), req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs: got length %d, want %d", len(resp.FSubs), len(req.FSubs))
	}
	for i := range req.FSubs {
		if resp.FSubs[i].FString != req.FSubs[i].FString {
			t.Errorf("FSubs[%d].FString: got %q, want %q", i, resp.FSubs[i].FString, req.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools: got length %d, want %d", len(resp.FBools), len(req.FBools))
	}
	for i := range req.FBools {
		if resp.FBools[i] != req.FBools[i] {
			t.Errorf("FBools[%d]: got %v, want %v", i, resp.FBools[i], req.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64: got %d, want %d", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S: got length %d, want %d", len(resp.FInt64S), len(req.FInt64S))
	}
	for i := range req.FInt64S {
		if resp.FInt64S[i] != req.FInt64S[i] {
			t.Errorf("FInt64S[%d]: got %d, want %d", i, resp.FInt64S[i], req.FInt64S[i])
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes: got %q, want %q", string(resp.FBytes), string(req.FBytes))
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess: got length %d, want %d", len(resp.FBytess), len(req.FBytess))
	}
	for i := range req.FBytess {
		if string(resp.FBytess[i]) != string(req.FBytess[i]) {
			t.Errorf("FBytess[%d]: got %q, want %q", i, string(resp.FBytess[i]), string(req.FBytess[i]))
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat: got %.3f, want %.3f", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats: got length %d, want %d", len(resp.FFloats), len(req.FFloats))
	}
	for i := range req.FFloats {
		if resp.FFloats[i] != req.FFloats[i] {
			t.Errorf("FFloats[%d]: got %.3f, want %.3f", i, resp.FFloats[i], req.FFloats[i])
		}
	}

	// Optionally, close send and check the stream closes
	if err := stream.CloseSend(); err != nil {
		t.Errorf("Error on close send: %v", err)
	}
}
