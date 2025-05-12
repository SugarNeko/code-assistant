package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createClient(t *testing.T) (grpcbin.GRPCBinClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	return grpcbin.NewGRPCBinClient(conn), conn
}

func TestDummyUnary_Positive(t *testing.T) {
	client, conn := createClient(t)
	defer conn.Close()

	req := &grpcbin.DummyMessage{
		FString:    "test",
		FStrings:   []string{"a", "b"},
		FInt32:     42,
		FInt32s:    []int32{1, 2},
		FEnum:      grpcbin.DummyMessage_ENUM_1,
		FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:       &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:      true,
		FBools:     []bool{true, false},
		FInt64:     1234567890,
		FInt64s:    []int64{987654321, 123456789},
		FBytes:     []byte("test bytes"),
		FBytess:    [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:     1.23,
		FFloats:    []float32{4.56, 7.89},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.GetFString() != req.GetFString() {
		t.Errorf("FString mismatch: got %v want %v", resp.GetFString(), req.GetFString())
	}
	if !reflect.DeepEqual(resp.GetFStrings(), req.GetFStrings()) {
		t.Errorf("FStrings mismatch: got %v want %v", resp.GetFStrings(), req.GetFStrings())
	}
	if resp.GetFInt32() != req.GetFInt32() {
		t.Errorf("FInt32 mismatch: got %v want %v", resp.GetFInt32(), req.GetFInt32())
	}
	if !reflect.DeepEqual(resp.GetFInt32s(), req.GetFInt32s()) {
		t.Errorf("FInt32s mismatch: got %v want %v", resp.GetFInt32s(), req.GetFInt32s())
	}
	if resp.GetFEnum() != req.GetFEnum() {
		t.Errorf("FEnum mismatch: got %v want %v", resp.GetFEnum(), req.GetFEnum())
	}
	if !reflect.DeepEqual(resp.GetFEnums(), req.GetFEnums()) {
		t.Errorf("FEnums mismatch: got %v want %v", resp.GetFEnums(), req.GetFEnums())
	}
	if !reflect.DeepEqual(resp.GetFSub(), req.GetFSub()) {
		t.Errorf("FSub mismatch: got %v want %v", resp.GetFSub(), req.GetFSub())
	}
	if !reflect.DeepEqual(resp.GetFSubs(), req.GetFSubs()) {
		t.Errorf("FSubs mismatch: got %v want %v", resp.GetFSubs(), req.GetFSubs())
	}
	if resp.GetFBool() != req.GetFBool() {
		t.Errorf("FBool mismatch: got %v want %v", resp.GetFBool(), req.GetFBool())
	}
	if !reflect.DeepEqual(resp.GetFBools(), req.GetFBools()) {
		t.Errorf("FBools mismatch: got %v want %v", resp.GetFBools(), req.GetFBools())
	}
	if resp.GetFInt64() != req.GetFInt64() {
		t.Errorf("FInt64 mismatch: got %v want %v", resp.GetFInt64(), req.GetFInt64())
	}
	if !reflect.DeepEqual(resp.GetFInt64s(), req.GetFInt64s()) {
		t.Errorf("FInt64s mismatch: got %v want %v", resp.GetFInt64s(), req.GetFInt64s())
	}
	if !reflect.DeepEqual(resp.GetFBytes(), req.GetFBytes()) {
		t.Errorf("FBytes mismatch: got %v want %v", resp.GetFBytes(), req.GetFBytes())
	}
	if !reflect.DeepEqual(resp.GetFBytess(), req.GetFBytess()) {
		t.Errorf("FBytess mismatch: got %v want %v", resp.GetFBytess(), req.GetFBytess())
	}
	if resp.GetFFloat() != req.GetFFloat() {
		t.Errorf("FFloat mismatch: got %v want %v", resp.GetFFloat(), req.GetFFloat())
	}
	if !reflect.DeepEqual(resp.GetFFloats(), req.GetFFloats()) {
		t.Errorf("FFloats mismatch: got %v want %v", resp.GetFFloats(), req.GetFFloats())
	}
}
