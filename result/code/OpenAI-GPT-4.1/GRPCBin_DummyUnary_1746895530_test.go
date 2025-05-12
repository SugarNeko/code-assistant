package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"one", "two"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "subfield"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    9876543210,
		FInt64S:   []int64{1, 2, 3, 4},
		FBytes:    []byte("bytes-test"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    1.23,
		FFloats:   []float32{3.14, 2.718},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary RPC failed: %v", err)
	}

	// Validate server echoed the same request back as response
	if resp.FString != req.FString {
		t.Errorf("f_string mismatch: got=%v want=%v", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("f_strings length mismatch: got=%d want=%d", len(resp.FStrings), len(req.FStrings))
	}
	for i := range req.FStrings {
		if resp.FStrings[i] != req.FStrings[i] {
			t.Errorf("f_strings[%d] mismatch: got=%v want=%v", i, resp.FStrings[i], req.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("f_int32 mismatch: got=%v want=%v", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("f_int32s length mismatch: got=%d want=%d", len(resp.FInt32S), len(req.FInt32S))
	}
	for i := range req.FInt32S {
		if resp.FInt32S[i] != req.FInt32S[i] {
			t.Errorf("f_int32s[%d] mismatch: got=%v want=%v", i, resp.FInt32S[i], req.FInt32S[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("f_enum mismatch: got=%v want=%v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("f_enums length mismatch: got=%d want=%d", len(resp.FEnums), len(req.FEnums))
	}
	for i := range req.FEnums {
		if resp.FEnums[i] != req.FEnums[i] {
			t.Errorf("f_enums[%d] mismatch: got=%v want=%v", i, resp.FEnums[i], req.FEnums[i])
		}
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("f_sub mismatch: got=%v want=%v", resp.FSub, req.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("f_subs length mismatch: got=%d want=%d", len(resp.FSubs), len(req.FSubs))
	}
	for i := range req.FSubs {
		if resp.FSubs[i].FString != req.FSubs[i].FString {
			t.Errorf("f_subs[%d].f_string mismatch: got=%v want=%v", i, resp.FSubs[i].FString, req.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("f_bool mismatch: got=%v want=%v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("f_bools length mismatch: got=%d want=%d", len(resp.FBools), len(req.FBools))
	}
	for i := range req.FBools {
		if resp.FBools[i] != req.FBools[i] {
			t.Errorf("f_bools[%d] mismatch: got=%v want=%v", i, resp.FBools[i], req.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("f_int64 mismatch: got=%v want=%v", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("f_int64s length mismatch: got=%d want=%d", len(resp.FInt64S), len(req.FInt64S))
	}
	for i := range req.FInt64S {
		if resp.FInt64S[i] != req.FInt64S[i] {
			t.Errorf("f_int64s[%d] mismatch: got=%v want=%v", i, resp.FInt64S[i], req.FInt64S[i])
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("f_bytes mismatch: got=%v want=%v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("f_bytess length mismatch: got=%d want=%d", len(resp.FBytess), len(req.FBytess))
	}
	for i := range req.FBytess {
		if string(resp.FBytess[i]) != string(req.FBytess[i]) {
			t.Errorf("f_bytess[%d] mismatch: got=%v want=%v", i, resp.FBytess[i], req.FBytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("f_float mismatch: got=%v want=%v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("f_floats length mismatch: got=%d want=%d", len(resp.FFloats), len(req.FFloats))
	}
	for i := range req.FFloats {
		if resp.FFloats[i] != req.FFloats[i] {
			t.Errorf("f_floats[%d] mismatch: got=%v want=%v", i, resp.FFloats[i], req.FFloats[i])
		}
	}
}
