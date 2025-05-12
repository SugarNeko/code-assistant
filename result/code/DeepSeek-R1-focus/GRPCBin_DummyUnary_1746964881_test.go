package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "test",
		FStrings:  []string{"a", "b"},
		FInt32:    123,
		FInt32s:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    456,
		FInt64s:   []int64{4, 5, 6},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

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
		for i, sub := range req.FSubs {
			if resp.FSubs[i].FString != sub.FString {
				t.Errorf("FSubs[%d].FString mismatch: got %q, want %q", i, resp.FSubs[i].FString, sub.FString)
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
	if !reflect.DeepEqual(resp.FBytes, req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
	} else {
		for i, b := range req.FBytess {
			if !reflect.DeepEqual(resp.FBytess[i], b) {
				t.Errorf("FBytess[%d] mismatch", i)
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
