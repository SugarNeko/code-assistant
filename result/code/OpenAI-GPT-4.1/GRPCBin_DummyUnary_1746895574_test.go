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
		FString:   "test-string",
		FStrings:  []string{"string1", "string2"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "nested"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    1234567890,
		FInt64S:   []int64{100, 200, 300},
		FBytes:    []byte("bytesdata"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{1.23, 4.56},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Client response validation (should echo the request)
	if resp.FString != req.FString {
		t.Errorf("got FString: %v, want: %v", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("got FStrings len: %v, want: %v", len(resp.FStrings), len(req.FStrings))
	}
	for i := range req.FStrings {
		if resp.FStrings[i] != req.FStrings[i] {
			t.Errorf("got FStrings[%d]: %v, want: %v", i, resp.FStrings[i], req.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("got FInt32: %v, want: %v", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("got FInt32S len: %v, want: %v", len(resp.FInt32S), len(req.FInt32S))
	}
	for i := range req.FInt32S {
		if resp.FInt32S[i] != req.FInt32S[i] {
			t.Errorf("got FInt32S[%d]: %v, want: %v", i, resp.FInt32S[i], req.FInt32S[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("got FEnum: %v, want: %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("got FEnums len: %v, want: %v", len(resp.FEnums), len(req.FEnums))
	}
	for i := range req.FEnums {
		if resp.FEnums[i] != req.FEnums[i] {
			t.Errorf("got FEnums[%d]: %v, want: %v", i, resp.FEnums[i], req.FEnums[i])
		}
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("got FSub: %v, want: %v", resp.FSub, req.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("got FSubs len: %v, want: %v", len(resp.FSubs), len(req.FSubs))
	}
	for i := range req.FSubs {
		if resp.FSubs[i].FString != req.FSubs[i].FString {
			t.Errorf("got FSubs[%d]: %v, want: %v", i, resp.FSubs[i].FString, req.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("got FBool: %v, want: %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("got FBools len: %v, want: %v", len(resp.FBools), len(req.FBools))
	}
	for i := range req.FBools {
		if resp.FBools[i] != req.FBools[i] {
			t.Errorf("got FBools[%d]: %v, want: %v", i, resp.FBools[i], req.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("got FInt64: %v, want: %v", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("got FInt64S len: %v, want: %v", len(resp.FInt64S), len(req.FInt64S))
	}
	for i := range req.FInt64S {
		if resp.FInt64S[i] != req.FInt64S[i] {
			t.Errorf("got FInt64S[%d]: %v, want: %v", i, resp.FInt64S[i], req.FInt64S[i])
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("got FBytes: %v, want: %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("got FBytess len: %v, want: %v", len(resp.FBytess), len(req.FBytess))
	}
	for i := range req.FBytess {
		if string(resp.FBytess[i]) != string(req.FBytess[i]) {
			t.Errorf("got FBytess[%d]: %v, want: %v", i, resp.FBytess[i], req.FBytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("got FFloat: %v, want: %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("got FFloats len: %v, want: %v", len(resp.FFloats), len(req.FFloats))
	}
	for i := range req.FFloats {
		if resp.FFloats[i] != req.FFloats[i] {
			t.Errorf("got FFloats[%d]: %v, want: %v", i, resp.FFloats[i], req.FFloats[i])
		}
	}
}
