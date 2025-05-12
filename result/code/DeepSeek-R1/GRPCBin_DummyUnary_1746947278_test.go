package grpcbin_test

import (
	"context"
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
		FStrings:  []string{"a", "b"},
		FInt32:    123,
		FInt32s:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-test"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    456,
		FInt64s:   []int64{4, 5, 6},
		FBytes:    []byte("test bytes"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.GetFString() != req.FString {
		t.Errorf("f_string mismatch: got %v, want %v", resp.GetFString(), req.FString)
	}

	if len(resp.GetFStrings()) != len(req.FStrings) {
		t.Errorf("f_strings length mismatch: got %d, want %d", len(resp.GetFStrings()), len(req.FStrings))
	} else {
		for i, v := range resp.GetFStrings() {
			if v != req.FStrings[i] {
				t.Errorf("f_strings[%d] mismatch: got %v, want %v", i, v, req.FStrings[i])
			}
		}
	}

	if resp.GetFInt32() != req.FInt32 {
		t.Errorf("f_int32 mismatch: got %v, want %v", resp.GetFInt32(), req.FInt32)
	}

	if len(resp.GetFInt32s()) != len(req.FInt32s) {
		t.Errorf("f_int32s length mismatch: got %d, want %d", len(resp.GetFInt32s()), len(req.FInt32s))
	} else {
		for i, v := range resp.GetFInt32s() {
			if v != req.FInt32s[i] {
				t.Errorf("f_int32s[%d] mismatch: got %v, want %v", i, v, req.FInt32s[i])
			}
		}
	}

	if resp.GetFEnum() != req.FEnum {
		t.Errorf("f_enum mismatch: got %v, want %v", resp.GetFEnum(), req.FEnum)
	}

	if len(resp.GetFEnums()) != len(req.FEnums) {
		t.Errorf("f_enums length mismatch: got %d, want %d", len(resp.GetFEnums()), len(req.FEnums()))
	} else {
		for i, v := range resp.GetFEnums() {
			if v != req.FEnums[i] {
				t.Errorf("f_enums[%d] mismatch: got %v, want %v", i, v, req.FEnums[i])
			}
		}
	}

	if resp.GetFSub().GetFString() != req.FSub.GetFString() {
		t.Errorf("f_sub.f_string mismatch: got %v, want %v", resp.GetFSub().GetFString(), req.FSub.GetFString())
	}

	if len(resp.GetFSubs()) != len(req.FSubs) {
		t.Errorf("f_subs length mismatch: got %d, want %d", len(resp.GetFSubs()), len(req.FSubs))
	} else {
		for i, v := range resp.GetFSubs() {
			if v.GetFString() != req.FSubs[i].GetFString() {
				t.Errorf("f_subs[%d].f_string mismatch: got %v, want %v", i, v.GetFString(), req.FSubs[i].GetFString())
			}
		}
	}

	if resp.GetFBool() != req.FBool {
		t.Errorf("f_bool mismatch: got %v, want %v", resp.GetFBool(), req.FBool)
	}

	if len(resp.GetFBools()) != len(req.FBools) {
		t.Errorf("f_bools length mismatch: got %d, want %d", len(resp.GetFBools()), len(req.FBools))
	} else {
		for i, v := range resp.GetFBools() {
			if v != req.FBools[i] {
				t.Errorf("f_bools[%d] mismatch: got %v, want %v", i, v, req.FBools[i])
			}
		}
	}

	if resp.GetFInt64() != req.FInt64 {
		t.Errorf("f_int64 mismatch: got %v, want %v", resp.GetFInt64(), req.FInt64)
	}

	if len(resp.GetFInt64s()) != len(req.FInt64s) {
		t.Errorf("f_int64s length mismatch: got %d, want %d", len(resp.GetFInt64s()), len(req.FInt64s))
	} else {
		for i, v := range resp.GetFInt64s() {
			if v != req.FInt64s[i] {
				t.Errorf("f_int64s[%d] mismatch: got %v, want %v", i, v, req.FInt64s[i])
			}
		}
	}

	if string(resp.GetFBytes()) != string(req.FBytes) {
		t.Errorf("f_bytes mismatch: got %v, want %v", resp.GetFBytes(), req.FBytes)
	}

	if len(resp.GetFBytess()) != len(req.FBytess) {
		t.Errorf("f_bytess length mismatch: got %d, want %d", len(resp.GetFBytess()), len(req.FBytess))
	} else {
		for i, v := range resp.GetFBytess() {
			if string(v) != string(req.FBytess[i]) {
				t.Errorf("f_bytess[%d] mismatch: got %v, want %v", i, v, req.FBytess[i])
			}
		}
	}

	if resp.GetFFloat() != req.FFloat {
		t.Errorf("f_float mismatch: got %v, want %v", resp.GetFFloat(), req.FFloat)
	}

	if len(resp.GetFFloats()) != len(req.FFloats) {
		t.Errorf("f_floats length mismatch: got %d, want %d", len(resp.GetFFloats()), len(req.FFloats))
	} else {
		for i, v := range resp.GetFFloats() {
			if v != req.FFloats[i] {
				t.Errorf("f_floats[%d] mismatch: got %v, want %v", i, v, req.FFloats[i])
			}
		}
	}
}
