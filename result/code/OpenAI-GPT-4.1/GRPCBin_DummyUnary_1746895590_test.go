package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to dial grpcb.in:9000: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub message"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "list sub 1"}, {FString: "list sub 2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    456789,
		FInt64S:   []int64{10, 20, 30},
		FBytes:    []byte{0x01, 0x02, 0x03},
		FBytess:   [][]byte{[]byte{0x04}, []byte{0x05}},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	// Validate server echoes the same DummyMessage
	if got, want := resp.FString, req.FString; got != want {
		t.Errorf("FString = %q; want %q", got, want)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length = %d; want %d", len(resp.FStrings), len(req.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 = %d; want %d", resp.FInt32, req.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum = %v; want %v", resp.FEnum, req.FEnum)
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString = %v; want %v", resp.FSub, req.FSub)
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool = %v; want %v", resp.FBool, req.FBool)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 = %d; want %d", resp.FInt64, req.FInt64)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes = %v; want %v", resp.FBytes, req.FBytes)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat = %v; want %v", resp.FFloat, req.FFloat)
	}
	// Optionally add further deep equality checks for slices if desired
}
