package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:    "test",
		FStrings:   []string{"a", "b"},
		FInt32:     42,
		FInt32S:    []int32{1, 2},
		FEnum:      grpcbin.DummyMessage_ENUM_1,
		FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:       &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:      true,
		FBools:     []bool{true, false},
		FInt64:     1234567890,
		FInt64S:    []int64{9876543210},
		FBytes:     []byte{0x01, 0x02},
		FBytess:    [][]byte{{0x03}, {0x04}},
		FFloat:     3.14,
		FFloats:    []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %v, want %v", resp.FString, req.FString)
	}

	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
	} else {
		for i := range resp.FStrings {
			if resp.FStrings[i] != req.FStrings[i] {
				t.Errorf("FStrings[%d] mismatch: got %v, want %v", i, resp.FStrings[i], req.FStrings[i])
			}
		}
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
	}

	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %d, want %d", len(resp.FInt32S), len(req.FInt32S))
	} else {
		for i := range resp.FInt32S {
			if resp.FInt32S[i] != req.FInt32S[i] {
				t.Errorf("FInt32S[%d] mismatch: got %v, want %v", i, resp.FInt32S[i], req.FInt32S[i])
			}
		}
	}

	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}

	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(resp.FEnums), len(req.FEnums))
	} else {
		for i := range resp.FEnums {
			if resp.FEnums[i] != req.FEnums[i] {
				t.Errorf("FEnums[%d] mismatch: got %v, want %v", i, resp.FEnums[i], req.FEnums[i])
			}
		}
	}

	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v, want %v", resp.FSub.FString, req.FSub.FString)
	}

	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
	} else {
		for i := range resp.FSubs {
			if resp.FSubs[i].FString != req.FSubs[i].FString {
				t.Errorf("FSubs[%d].FString mismatch: got %v, want %v", i, resp.FSubs[i].FString, req.FSubs[i].FString)
			}
		}
	}

	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}

	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(resp.FBools), len(req.FBools))
	} else {
		for i := range resp.FBools {
			if resp.FBools[i] != req.FBools[i] {
				t.Errorf("FBools[%d] mismatch: got %v, want %v", i, resp.FBools[i], req.FBools[i])
			}
		}
	}

	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, req.FInt64)
	}

	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %d, want %d", len(resp.FInt64S), len(req.FInt64S))
	} else {
		for i := range resp.FInt64S {
			if resp.FInt64S[i] != req.FInt64S[i] {
				t.Errorf("FInt64S[%d] mismatch: got %v, want %v", i, resp.FInt64S[i], req.FInt64S[i])
			}
		}
	}

	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
	}

	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
	} else {
		for i := range resp.FBytess {
			if string(resp.FBytess[i]) != string(req.FBytess[i]) {
				t.Errorf("FBytess[%d] mismatch: got %v, want %v", i, resp.FBytess[i], req.FBytess[i])
			}
		}
	}

	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, req.FFloat)
	}

	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(resp.FFloats), len(req.FFloats))
	} else {
		for i := range resp.FFloats {
			if resp.FFloats[i] != req.FFloats[i] {
				t.Errorf("FFloats[%d] mismatch: got %v, want %v", i, resp.FFloats[i], req.FFloats[i])
			}
		}
	}
}
