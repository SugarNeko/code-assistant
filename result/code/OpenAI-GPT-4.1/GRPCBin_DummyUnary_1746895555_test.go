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
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"foo", "bar"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_1},
		FSub:      &pb.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub-1"}, {FString: "sub-2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    456789,
		FInt64S:   []int64{111, 222, 333},
		FBytes:    []byte("abc"),
		FBytess:   [][]byte{[]byte("hello"), []byte("world")},
		FFloat:    1.23,
		FFloats:   []float32{4.56, 7.89},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	if resp == nil {
		t.Fatalf("response is nil")
	}

	// response validation
	if resp.FString != req.FString {
		t.Errorf("unexpected FString. got=%q want=%q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("unexpected FStrings length. got=%d want=%d", len(resp.FStrings), len(req.FStrings))
	}
	for i := range req.FStrings {
		if resp.FStrings[i] != req.FStrings[i] {
			t.Errorf("unexpected FStrings[%d]. got=%q want=%q", i, resp.FStrings[i], req.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("unexpected FInt32. got=%d want=%d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("unexpected FInt32S length. got=%d want=%d", len(resp.FInt32S), len(req.FInt32S))
	}
	for i := range req.FInt32S {
		if resp.FInt32S[i] != req.FInt32S[i] {
			t.Errorf("unexpected FInt32S[%d]. got=%d want=%d", i, resp.FInt32S[i], req.FInt32S[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("unexpected FEnum. got=%v want=%v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("unexpected FEnums length. got=%d want=%d", len(resp.FEnums), len(req.FEnums))
	}
	for i := range req.FEnums {
		if resp.FEnums[i] != req.FEnums[i] {
			t.Errorf("unexpected FEnums[%d]. got=%v want=%v", i, resp.FEnums[i], req.FEnums[i])
		}
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("unexpected FSub.FString. got=%q want=%q", resp.FSub.FString, req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("unexpected FSubs length. got=%d want=%d", len(resp.FSubs), len(req.FSubs))
	}
	for i := range req.FSubs {
		if resp.FSubs[i].FString != req.FSubs[i].FString {
			t.Errorf("unexpected FSubs[%d].FString. got=%q want=%q", i, resp.FSubs[i].FString, req.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("unexpected FBool. got=%v want=%v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("unexpected FBools length. got=%d want=%d", len(resp.FBools), len(req.FBools))
	}
	for i := range req.FBools {
		if resp.FBools[i] != req.FBools[i] {
			t.Errorf("unexpected FBools[%d]. got=%v want=%v", i, resp.FBools[i], req.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("unexpected FInt64. got=%d want=%d", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("unexpected FInt64S length. got=%d want=%d", len(resp.FInt64S), len(req.FInt64S))
	}
	for i := range req.FInt64S {
		if resp.FInt64S[i] != req.FInt64S[i] {
			t.Errorf("unexpected FInt64S[%d]. got=%d want=%d", i, resp.FInt64S[i], req.FInt64S[i])
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("unexpected FBytes. got=%q want=%q", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("unexpected FBytess length. got=%d want=%d", len(resp.FBytess), len(req.FBytess))
	}
	for i := range req.FBytess {
		if string(resp.FBytess[i]) != string(req.FBytess[i]) {
			t.Errorf("unexpected FBytess[%d]. got=%q want=%q", i, resp.FBytess[i], req.FBytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("unexpected FFloat. got=%f want=%f", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("unexpected FFloats length. got=%d want=%d", len(resp.FFloats), len(req.FFloats))
	}
	for i := range req.FFloats {
		if resp.FFloats[i] != req.FFloats[i] {
			t.Errorf("unexpected FFloats[%d]. got=%f want=%f", i, resp.FFloats[i], req.FFloats[i])
		}
	}
}
