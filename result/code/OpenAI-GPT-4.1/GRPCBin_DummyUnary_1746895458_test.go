package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "substring"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "one"}, {FString: "two"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    456789012,
		FInt64S:   []int64{1001, 1002},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("foo"), []byte("bar")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary returned error: %v", err)
	}

	// Client-side response validation (echo should be same as sent)
	if resp.FString != req.FString {
		t.Errorf("FString: got %q, want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings: got %d elements, want %d", len(resp.FStrings), len(req.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32: got %d, want %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S: got %v, want %v", resp.FInt32S, req.FInt32S)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums: got %v, want %v", resp.FEnums, req.FEnums)
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub: got %+v, want %+v", resp.FSub, req.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs: got %+v, want %+v", resp.FSubs, req.FSubs)
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools: got %v, want %v", resp.FBools, req.FBools)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64: got %d, want %d", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S: got %v, want %v", resp.FInt64S, req.FInt64S)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess: got %v, want %v", resp.FBytess, req.FBytess)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat: got %v, want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats: got %v, want %v", resp.FFloats, req.FFloats)
	}
}
