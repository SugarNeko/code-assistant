package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "sub-string",
		},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{FString: "sublist-1"},
			{FString: "sublist-2"},
		},
		FBool:   true,
		FBools:  []bool{false, true},
		FInt64:  9001,
		FInt64S: []int64{11, 12, 13},
		FBytes:  []byte("hello-bytes"),
		FBytess: [][]byte{[]byte("a"), []byte("b")},
		FFloat:  3.14,
		FFloats: []float32{1.1, 2.2, 3.3},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	// Validate response equals the request
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %v, want %v", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %v, want %v", len(resp.FStrings), len(req.FStrings))
	}
	for i, v := range req.FStrings {
		if resp.FStrings[i] != v {
			t.Errorf("FStrings[%d] mismatch: got %v, want %v", i, resp.FStrings[i], v)
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
			t.Errorf("FInt32S[%d] mismatch: got %v, want %v", i, resp.FInt32S[i], v)
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
	if (resp.FSub != nil && req.FSub == nil) || (resp.FSub == nil && req.FSub != nil) {
		t.Errorf("FSub presence mismatch: got %v, want %v", resp.FSub, req.FSub)
	} else if resp.FSub != nil && req.FSub != nil && resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v, want %v", resp.FSub.FString, req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
	}
	for i, v := range req.FSubs {
		if resp.FSubs[i].FString != v.FString {
			t.Errorf("FSubs[%d].FString mismatch: got %v, want %v", i, resp.FSubs[i].FString, v.FString)
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
		t.Errorf("FBytes mismatch: got %v, want %v", string(resp.FBytes), string(req.FBytes))
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
	}
	for i, v := range req.FBytess {
		if string(resp.FBytess[i]) != string(v) {
			t.Errorf("FBytess[%d] mismatch: got %v, want %v", i, string(resp.FBytess[i]), string(v))
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
}
