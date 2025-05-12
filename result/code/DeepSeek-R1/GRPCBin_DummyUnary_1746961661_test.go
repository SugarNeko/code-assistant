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
		t.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"a", "b", "c"},
		FInt32:    42,
		FInt32S:   []int32{10, 20, 30},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    123456789,
		FInt64S:   []int64{987654321, 123123123},
		FBytes:    []byte{0x01, 0x02, 0x03},
		FBytess:   [][]byte{{0x04, 0x05}, {0x06}},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("RPC failed: %v", err)
	}

	if !reflect.DeepEqual(resp, req) {
		t.Error("Response does not match request:")
		if resp.FString != req.FString {
			t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
		}
		if !reflect.DeepEqual(resp.FStrings, req.FStrings) {
			t.Errorf("FStrings mismatch: got %v, want %v", resp.FStrings, req.FStrings)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
		}
		if !reflect.DeepEqual(resp.FInt32S, req.FInt32S) {
			t.Errorf("FInt32s mismatch: got %v, want %v", resp.FInt32S, req.FInt32S)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if !reflect.DeepEqual(resp.FEnums, req.FEnums) {
			t.Errorf("FEnums mismatch: got %v, want %v", resp.FEnums, req.FEnums)
		}
		if !reflect.DeepEqual(resp.FSub, req.FSub) {
			t.Errorf("FSub mismatch: got %v, want %v", resp.FSub, req.FSub)
		}
		if !reflect.DeepEqual(resp.FSubs, req.FSubs) {
			t.Errorf("FSubs mismatch: got %v, want %v", resp.FSubs, req.FSubs)
		}
		if resp.FBool != req.FBool {
			t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
		}
		if !reflect.DeepEqual(resp.FBools, req.FBools) {
			t.Errorf("FBools mismatch: got %v, want %v", resp.FBools, req.FBools)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
		}
		if !reflect.DeepEqual(resp.FInt64S, req.FInt64S) {
			t.Errorf("FInt64s mismatch: got %v, want %v", resp.FInt64S, req.FInt64S)
		}
		if !reflect.DeepEqual(resp.FBytes, req.FBytes) {
			t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
		}
		if !reflect.DeepEqual(resp.FBytess, req.FBytess) {
			t.Errorf("FBytess mismatch: got %v, want %v", resp.FBytess, req.FBytess)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("FFloat mismatch: got %f, want %f", resp.FFloat, req.FFloat)
		}
		if !reflect.DeepEqual(resp.FFloats, req.FFloats) {
			t.Errorf("FFloats mismatch: got %v, want %v", resp.FFloats, req.FFloats)
		}
	}
}
