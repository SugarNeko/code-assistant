package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:    "test-string",
		FStrings:   []string{"foo", "bar"},
		FInt32:     123,
		FInt32S:    []int32{7, 8, 9},
		FEnum:      grpcbin.DummyMessage_ENUM_1,
		FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSub:       &grpcbin.DummyMessage_Sub{FString: "subfield"},
		FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "subfield1"}, {FString: "subfield2"}},
		FBool:      true,
		FBools:     []bool{true, false},
		FInt64:     456789,
		FInt64S:    []int64{111, 222},
		FBytes:     []byte("bytestring"),
		FBytess:    [][]byte{[]byte("first"), []byte("second")},
		FFloat:     1.23,
		FFloats:    []float32{3.14, 2.71},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Client response validation and server echo validation
	if resp.FString != req.FString {
		t.Errorf("FString: got %v want %v", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings: got %v want %v", resp.FStrings, req.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32: got %v want %v", resp.FInt32, req.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum: got %v want %v", resp.FEnum, req.FEnum)
	}
	if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub: got %v want %v", resp.FSub, req.FSub)
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool: got %v want %v", resp.FBool, req.FBool)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64: got %v want %v", resp.FInt64, req.FInt64)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes: got %v want %v", resp.FBytes, req.FBytes)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat: got %v want %v", resp.FFloat, req.FFloat)
	}
	// More detailed checks for slices can be added as necessary
}
