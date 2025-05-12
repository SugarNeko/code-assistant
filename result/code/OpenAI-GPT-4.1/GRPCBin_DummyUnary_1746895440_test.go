package grpcbin_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestGRPCBin_DummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to grpcb.in:9000: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"string-1", "string-2"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub-1"}, {FString: "sub-2"}},
		FBool:    true,
		FBools:   []bool{true, false, true},
		FInt64:   1234567890,
		FInt64S:  []int64{100, 200, 300},
		FBytes:   []byte("hello-bytes"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Response validation: check that the response matches the sent request fields
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: want %q, got %q", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: want %d, got %d", len(req.FStrings), len(resp.FStrings))
	} else {
		for i, v := range req.FStrings {
			if resp.FStrings[i] != v {
				t.Errorf("FStrings[%d] mismatch: want %q, got %q", i, v, resp.FStrings[i])
			}
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: want %d, got %d", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length mismatch: want %d, got %d", len(req.FInt32S), len(resp.FInt32S))
	} else {
		for i, v := range req.FInt32S {
			if resp.FInt32S[i] != v {
				t.Errorf("FInt32S[%d] mismatch: want %d, got %d", i, v, resp.FInt32S[i])
			}
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: want %v, got %v", req.FEnum, resp.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: want %d, got %d", len(req.FEnums), len(resp.FEnums))
	} else {
		for i, v := range req.FEnums {
			if resp.FEnums[i] != v {
				t.Errorf("FEnums[%d] mismatch: want %v, got %v", i, v, resp.FEnums[i])
			}
		}
	}
	if req.FSub != nil && resp.FSub != nil {
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("FSub.FString mismatch: want %q, got %q", req.FSub.FString, resp.FSub.FString)
		}
	} else if !(req.FSub == nil && resp.FSub == nil) {
		t.Errorf("FSub nil mismatch: want %v, got %v", req.FSub, resp.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: want %d, got %d", len(req.FSubs), len(resp.FSubs))
	} else {
		for i, v := range req.FSubs {
			if resp.FSubs[i].FString != v.FString {
				t.Errorf("FSubs[%d].FString mismatch: want %q, got %q", i, v.FString, resp.FSubs[i].FString)
			}
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: want %v, got %v", req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: want %d, got %d", len(req.FBools), len(resp.FBools))
	} else {
		for i, v := range req.FBools {
			if resp.FBools[i] != v {
				t.Errorf("FBools[%d] mismatch: want %v, got %v", i, v, resp.FBools[i])
			}
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: want %d, got %d", req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: want %d, got %d", len(req.FInt64S), len(resp.FInt64S))
	} else {
		for i, v := range req.FInt64S {
			if resp.FInt64S[i] != v {
				t.Errorf("FInt64S[%d] mismatch: want %d, got %d", i, v, resp.FInt64S[i])
			}
		}
	}
	if !bytes.Equal(resp.FBytes, req.FBytes) {
		t.Errorf("FBytes mismatch: want %v, got %v", req.FBytes, resp.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: want %d, got %d", len(req.FBytess), len(resp.FBytess))
	} else {
		for i, v := range req.FBytess {
			if !bytes.Equal(resp.FBytess[i], v) {
				t.Errorf("FBytess[%d] mismatch: want %v, got %v", i, v, resp.FBytess[i])
			}
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: want %v, got %v", req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: want %d, got %d", len(req.FFloats), len(resp.FFloats))
	} else {
		for i, v := range req.FFloats {
			if resp.FFloats[i] != v {
				t.Errorf("FFloats[%d] mismatch: want %v, got %v", i, v, resp.FFloats[i])
			}
		}
	}
}
