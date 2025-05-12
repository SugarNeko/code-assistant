package grpcbin_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnaryPositive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"a", "b"},
		FInt32:   123,
		FInt32s:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{FString: "sub1"},
			{FString: "sub2"},
		},
		FBool:   true,
		FBools:  []bool{true, false},
		FInt64:  456,
		FInt64s: []int64{4, 5, 6},
		FBytes:  []byte("bytes"),
		FBytess: [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:  3.14,
		FFloats: []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary RPC failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
	}

	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
	} else {
		for i := range resp.FStrings {
			if resp.FStrings[i] != req.FStrings[i] {
				t.Errorf("FStrings[%d] mismatch: got %q, want %q", i, resp.FStrings[i], req.FStrings[i])
			}
		}
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
	}

	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("FInt32s length mismatch: got %d, want %d", len(resp.FInt32s), len(req.FInt32s))
	} else {
		for i := range resp.FInt32s {
			if resp.FInt32s[i] != req.FInt32s[i] {
				t.Errorf("FInt32s[%d] mismatch: got %d, want %d", i, resp.FInt32s[i], req.FInt32s[i])
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
		t.Errorf("FSub.FString mismatch: got %q, want %q", resp.FSub.FString, req.FSub.FString)
	}

	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
	} else {
		for i := range resp.FSubs {
			if resp.FSubs[i].FString != req.FSubs[i].FString {
				t.Errorf("FSubs[%d].FString mismatch: got %q, want %q", i, resp.FSubs[i].FString, req.FSubs[i].FString)
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
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
	}

	if len(resp.FInt64s) != len(req.FInt64s) {
		t.Errorf("FInt64s length mismatch: got %d, want %d", len(resp.FInt64s), len(req.FInt64s))
	} else {
		for i := range resp.FInt64s {
			if resp.FInt64s[i] != req.FInt64s[i] {
				t.Errorf("FInt64s[%d] mismatch: got %d, want %d", i, resp.FInt64s[i], req.FInt64s[i])
			}
		}
	}

	if !bytes.Equal(resp.FBytes, req.FBytes) {
		t.Errorf("FBytes mismatch: got %q, want %q", resp.FBytes, req.FBytes)
	}

	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
	} else {
		for i := range resp.FBytess {
			if !bytes.Equal(resp.FBytess[i], req.FBytess[i]) {
				t.Errorf("FBytess[%d] mismatch: got %q, want %q", i, resp.FBytess[i], req.FBytess[i])
			}
		}
	}

	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %f, want %f", resp.FFloat, req.FFloat)
	}

	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(resp.FFloats), len(req.FFloats))
	} else {
		for i := range resp.FFloats {
			if resp.FFloats[i] != req.FFloats[i] {
				t.Errorf("FFloats[%d] mismatch: got %f, want %f", i, resp.FFloats[i], req.FFloats[i])
			}
		}
	}
}
