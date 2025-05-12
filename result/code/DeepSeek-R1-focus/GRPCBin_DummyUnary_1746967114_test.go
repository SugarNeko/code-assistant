package grpcbin_test

import (
	"context"
	"testing"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "test",
		FStrings:  []string{"a", "b"},
		FInt32:    123,
		FInt32S:   []int32{1, 2},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    456,
		FInt64S:   []int64{3, 4},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    1.23,
		FFloats:   []float32{5.6, 7.8},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("RPC failed: %v", err)
	}

	if res.FString != req.FString {
		t.Errorf("FString mismatch: got %v, want %v", res.FString, req.FString)
	}
	if len(res.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(res.FStrings), len(req.FStrings))
	}
	if res.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", res.FInt32, req.FInt32)
	}
	if len(res.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32s length mismatch: got %d, want %d", len(res.FInt32S), len(req.FInt32S))
	}
	if res.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", res.FEnum, req.FEnum)
	}
	if len(res.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(res.FEnums), len(req.FEnums))
	}
	if res.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v, want %v", res.FSub.FString, req.FSub.FString)
	}
	if len(res.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(res.FSubs), len(req.FSubs))
	}
	if res.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", res.FBool, req.FBool)
	}
	if len(res.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(res.FBools), len(req.FBools))
	}
	if res.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", res.FInt64, req.FInt64)
	}
	if len(res.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64s length mismatch: got %d, want %d", len(res.FInt64S), len(req.FInt64S))
	}
	if string(res.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", res.FBytes, req.FBytes)
	}
	if len(res.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(res.FBytess), len(req.FBytess))
	}
	if res.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", res.FFloat, req.FFloat)
	}
	if len(res.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(res.FFloats), len(req.FFloats))
	}
}
