package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "hello",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_2,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "one"}, {FString: "two"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   123456789,
		FInt64S:  []int64{9, 8, 7},
		FBytes:   []byte("bytes-value"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   1.23,
		FFloats:  []float32{4.5, 6.7},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Client response validation
	if resp.FString != req.FString {
		t.Errorf("FString: got %q, want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings: got len %d, want %d", len(resp.FStrings), len(req.FStrings))
	}
	for i, v := range req.FStrings {
		if resp.FStrings[i] != v {
			t.Errorf("FStrings[%d]: got %q, want %q", i, resp.FStrings[i], v)
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32: got %d, want %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S: got len %d, want %d", len(resp.FInt32S), len(req.FInt32S))
	}
	for i, v := range req.FInt32S {
		if resp.FInt32S[i] != v {
			t.Errorf("FInt32S[%d]: got %d, want %d", i, resp.FInt32S[i], v)
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum: got %d, want %d", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums: got len %d, want %d", len(resp.FEnums), len(req.FEnums))
	}
	for i, v := range req.FEnums {
		if resp.FEnums[i] != v {
			t.Errorf("FEnums[%d]: got %d, want %d", i, resp.FEnums[i], v)
		}
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString: got %q, want %q", resp.FSub.FString, req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs: got len %d, want %d", len(resp.FSubs), len(req.FSubs))
	}
	for i, v := range req.FSubs {
		if resp.FSubs[i] == nil || resp.FSubs[i].FString != v.FString {
			t.Errorf("FSubs[%d].FString: got %q, want %q", i, resp.FSubs[i].FString, v.FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools: got len %d, want %d", len(resp.FBools), len(req.FBools))
	}
	for i, v := range req.FBools {
		if resp.FBools[i] != v {
			t.Errorf("FBools[%d]: got %v, want %v", i, resp.FBools[i], v)
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64: got %d, want %d", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S: got len %d, want %d", len(resp.FInt64S), len(req.FInt64S))
	}
	for i, v := range req.FInt64S {
		if resp.FInt64S[i] != v {
			t.Errorf("FInt64S[%d]: got %d, want %d", i, resp.FInt64S[i], v)
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess: got len %d, want %d", len(resp.FBytess), len(req.FBytess))
	}
	for i, v := range req.FBytess {
		if string(resp.FBytess[i]) != string(v) {
			t.Errorf("FBytess[%d]: got %v, want %v", i, resp.FBytess[i], v)
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat: got %v, want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats: got len %d, want %d", len(resp.FFloats), len(req.FFloats))
	}
	for i, v := range req.FFloats {
		if resp.FFloats[i] != v {
			t.Errorf("FFloats[%d]: got %v, want %v", i, resp.FFloats[i], v)
		}
	}
}
