package grpcbin_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "test",
		FStrings:  []string{"a", "b", "c"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub_test"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    456,
		FInt64S:   []int64{4, 5, 6},
		FBytes:    []byte{0x01, 0x02},
		FBytess:   [][]byte{{0x03}, {0x04}},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	res, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if res.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", res.FString, req.FString)
	}
	if !reflect.DeepEqual(res.FStrings, req.FStrings) {
		t.Errorf("FStrings mismatch: got %v, want %v", res.FStrings, req.FStrings)
	}
	if res.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", res.FInt32, req.FInt32)
	}
	if !reflect.DeepEqual(res.FInt32S, req.FInt32S) {
		t.Errorf("FInt32S mismatch: got %v, want %v", res.FInt32S, req.FInt32S)
	}
	if res.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", res.FEnum, req.FEnum)
	}
	if !reflect.DeepEqual(res.FEnums, req.FEnums) {
		t.Errorf("FEnums mismatch: got %v, want %v", res.FEnums, req.FEnums)
	}
	if !reflect.DeepEqual(res.FSub, req.FSub) {
		t.Errorf("FSub mismatch: got %v, want %v", res.FSub, req.FSub)
	}
	if !reflect.DeepEqual(res.FSubs, req.FSubs) {
		t.Errorf("FSubs mismatch: got %v, want %v", res.FSubs, req.FSubs)
	}
	if res.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %t, want %t", res.FBool, req.FBool)
	}
	if !reflect.DeepEqual(res.FBools, req.FBools) {
		t.Errorf("FBools mismatch: got %v, want %v", res.FBools, req.FBools)
	}
	if res.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", res.FInt64, req.FInt64)
	}
	if !reflect.DeepEqual(res.FInt64S, req.FInt64S) {
		t.Errorf("FInt64S mismatch: got %v, want %v", res.FInt64S, req.FInt64S)
	}
	if !reflect.DeepEqual(res.FBytes, req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", res.FBytes, req.FBytes)
	}
	if !reflect.DeepEqual(res.FBytess, req.FBytess) {
		t.Errorf("FBytess mismatch: got %v, want %v", res.FBytess, req.FBytess)
	}
	if res.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %f, want %f", res.FFloat, req.FFloat)
	}
	if !reflect.DeepEqual(res.FFloats, req.FFloats) {
		t.Errorf("FFloats mismatch: got %v, want %v", res.FFloats, req.FFloats)
	}
}
