package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub_string"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    1001,
		FInt64S:   []int64{999, 1000},
		FBytes:    []byte("hello"),
		FBytess:   [][]byte{[]byte("foo"), []byte("bar")},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 0.577},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Validate that the response matches the request fields
	if resp.FString != req.FString {
		t.Errorf("FString want %v, got %v", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length want %d, got %d", len(req.FStrings), len(resp.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 want %d, got %d", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length want %d, got %d", len(req.FInt32S), len(resp.FInt32S))
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum want %v, got %v", req.FEnum, resp.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length want %d, got %d", len(req.FEnums), len(resp.FEnums))
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub want %v, got %v", req.FSub.FString, resp.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length want %d, got %d", len(req.FSubs), len(resp.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool want %v, got %v", req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length want %d, got %d", len(req.FBools), len(resp.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 want %d, got %d", req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length want %d, got %d", len(req.FInt64S), len(resp.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes want %v, got %v", req.FBytes, resp.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length want %d, got %d", len(req.FBytess), len(resp.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat want %v, got %v", req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length want %d, got %d", len(req.FFloats), len(resp.FFloats))
	}
}
