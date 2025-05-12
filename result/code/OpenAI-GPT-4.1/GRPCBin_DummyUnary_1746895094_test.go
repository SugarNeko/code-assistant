package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
	// Connect to the running gRPCBin server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct a typical (positive) DummyMessage request
	req := &grpcbin.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"val1", "val2"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub-str1"}, {FString: "sub-str2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    987654321,
		FInt64S:   []int64{1001, 1002},
		FBytes:    []byte("bytes-data"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{0.1, 0.2, 0.3},
	}

	// Make the DummyUnary call
	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	// Client response validation: check if response matches the request (echo)
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
	}
	for i := range req.FStrings {
		if resp.FStrings[i] != req.FStrings[i] {
			t.Errorf("FStrings[%d] mismatch: got %q, want %q", i, resp.FStrings[i], req.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %d, want %d", len(resp.FInt32S), len(req.FInt32S))
	}
	for i := range req.FInt32S {
		if resp.FInt32S[i] != req.FInt32S[i] {
			t.Errorf("FInt32S[%d] mismatch: got %d, want %d", i, resp.FInt32S[i], req.FInt32S[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(resp.FEnums), len(req.FEnums))
	}
	for i := range req.FEnums {
		if resp.FEnums[i] != req.FEnums[i] {
			t.Errorf("FEnums[%d] mismatch: got %v, want %v", i, resp.FEnums[i], req.FEnums[i])
		}
	}
	if req.FSub != nil && (resp.FSub == nil || resp.FSub.FString != req.FSub.FString) {
		t.Errorf("FSub mismatch: got %+v, want %+v", resp.FSub, req.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
	}
	for i := range req.FSubs {
		if resp.FSubs[i].FString != req.FSubs[i].FString {
			t.Errorf("FSubs[%d].FString mismatch: got %q, want %q", i, resp.FSubs[i].FString, req.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(resp.FBools), len(req.FBools))
	}
	for i := range req.FBools {
		if resp.FBools[i] != req.FBools[i] {
			t.Errorf("FBools[%d] mismatch: got %v, want %v", i, resp.FBools[i], req.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %d, want %d", len(resp.FInt64S), len(req.FInt64S))
	}
	for i := range req.FInt64S {
		if resp.FInt64S[i] != req.FInt64S[i] {
			t.Errorf("FInt64S[%d] mismatch: got %d, want %d", i, resp.FInt64S[i], req.FInt64S[i])
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %q, want %q", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
	}
	for i := range req.FBytess {
		if string(resp.FBytess[i]) != string(req.FBytess[i]) {
			t.Errorf("FBytess[%d] mismatch: got %q, want %q", i, resp.FBytess[i], req.FBytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %f, want %f", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(resp.FFloats), len(req.FFloats))
	}
	for i := range req.FFloats {
		if resp.FFloats[i] != req.FFloats[i] {
			t.Errorf("FFloats[%d] mismatch: got %f, want %f", i, resp.FFloats[i], req.FFloats[i])
		}
	}
}
