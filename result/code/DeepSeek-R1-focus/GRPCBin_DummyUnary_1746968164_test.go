package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"code-assistant/proto/grpcbin"
)

func TestDummyUnary_PositiveCase(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"one", "two"},
		FInt32:   42,
		FInt32s:  []int32{10, 20},
		FEnum:    grpcbin.DummyMessage_ENUM_2,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_0},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   9876543210,
		FInt64s:  []int64{100, 200},
		FBytes:   []byte{0x01, 0x02},
		FBytess:  [][]byte{{0x03}, {0x04}},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 9.9},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("RPC failed: %v", err)
	}

	if !reflect.DeepEqual(resp, req) {
		t.Error("Response doesn't match request:")
		if resp.FString != req.FString {
			t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
		}
		if !reflect.DeepEqual(resp.FStrings, req.FStrings) {
			t.Errorf("FStrings mismatch: got %v, want %v", resp.FStrings, req.FStrings)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
		}
		if !reflect.DeepEqual(resp.FInt32s, req.FInt32s) {
			t.Errorf("FInt32s mismatch: got %v, want %v", resp.FInt32s, req.FInt32s)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if !reflect.DeepEqual(resp.FEnums, req.FEnums) {
			t.Errorf("FEnums mismatch: got %v, want %v", resp.FEnums, req.FEnums)
		}
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("FSub.FString mismatch: got %q, want %q", resp.FSub.FString, req.FSub.FString)
		}
		if len(resp.FSubs) != len(req.FSubs) {
			t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
		} else {
			for i := range resp.FSubs {
				if resp.FSubs[i].FString != req.FSubs[i].FString {
					t.Errorf("FSubs[%d] mismatch: got %q, want %q", i, resp.FSubs[i].FString, req.FSubs[i].FString)
				}
			}
		}
		if resp.FBool != req.FBool {
			t.Errorf("FBool mismatch: got %t, want %t", resp.FBool, req.FBool)
		}
		if !reflect.DeepEqual(resp.FBools, req.FBools) {
			t.Errorf("FBools mismatch: got %v, want %v", resp.FBools, req.FBools)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
		}
		if !reflect.DeepEqual(resp.FInt64s, req.FInt64s) {
			t.Errorf("FInt64s mismatch: got %v, want %v", resp.FInt64s, req.FInt64s)
		}
		if !bytesEqual(resp.FBytes, req.FBytes) {
			t.Errorf("FBytes mismatch: got %x, want %x", resp.FBytes, req.FBytes)
		}
		if len(resp.FBytess) != len(req.FBytess) {
			t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
		} else {
			for i := range resp.FBytess {
				if !bytesEqual(resp.FBytess[i], req.FBytess[i]) {
					t.Errorf("FBytess[%d] mismatch: got %x, want %x", i, resp.FBytess[i], req.FBytess[i])
				}
			}
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("FFloat mismatch: got %f, want %f", resp.FFloat, req.FFloat)
		}
		if !reflect.DeepEqual(resp.FFloats, req.FFloats) {
			t.Errorf("FFloats mismatch: got %v, want %v", resp.FFloats, req.FFloats)
		}
	}
}

func bytesEqual(a, b []byte) bool {
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
