package grpcbin

import (
	"context"
	"testing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPCBinDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	// Positive testing
	req := &DummyMessage{
		FString: "test",
		FStrings: []string{"test1", "test2"},
		FInt32: 123,
		FInt32s: []int32{1, 2, 3},
		Enum:     Enum_ENUM_1,
		FEnums:   []Enum{Enum_ENUM_0, Enum_ENUM_1},
		FSub:     &Sub{FString: "sub"},
		FSubs: []*Sub{
			{FString: "sub1"},
			{FString: "sub2"},
		},
		FBool:   true,
		FBools:  []bool{true, false},
		FInt64:  123456,
		FInt64s: []int64{1, 2, 3},
		FBytes:  []byte("test"),
		FBytess: [][]byte{[]byte("test1"), []byte("test2")},
		FFloat:  12.34,
		FFloats: []float32{1.2, 3.4},
	}
	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("test failed: %v", err)
	}

	// Client response validation
	if resp.FString != req.FString {
		t.Errorf("response string mismatch: %s != %s", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("response strings length mismatch: %d != %d", len(resp.FStrings), len(req.FStrings))
	}
	for i, s := range resp.FStrings {
		if s != req.FStrings[i] {
			t.Errorf("response strings mismatch at index %d: %s != %s", i, s, req.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("response int32 mismatch: %d != %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("response int32s length mismatch: %d != %d", len(resp.FInt32s), len(req.FInt32s))
	}
	for i, n := range resp.FInt32s {
		if n != req.FInt32s[i] {
			t.Errorf("response int32s mismatch at index %d: %d != %d", i, n, req.FInt32s[i])
		}
	}
	if resp.Enum != req.Enum {
		t.Errorf("response enum mismatch: %d != %d", resp.Enum, req.Enum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("response enums length mismatch: %d != %d", len(resp.FEnums), len(req.FEnums))
	}
	for i, e := range resp.FEnums {
		if e != req.FEnums[i] {
			t.Errorf("response enums mismatch at index %d: %d != %d", i, e, req.FEnums[i])
		}
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("response sub string mismatch: %s != %s", resp.FSub.FString, req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("response subs length mismatch: %d != %d", len(resp.FSubs), len(req.FSubs))
	}
	for i, s := range resp.FSubs {
		if s.FString != req.FSubs[i].FString {
			t.Errorf("response subs string mismatch at index %d: %s != %s", i, s.FString, req.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("response bool mismatch: %t != %t", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("response bools length mismatch: %d != %d", len(resp.FBools), len(req.FBools))
	}
	for i, b := range resp.FBools {
		if b != req.FBools[i] {
			t.Errorf("response bools mismatch at index %d: %t != %t", i, b, req.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("response int64 mismatch: %d != %d", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64s) != len(req.FInt64s) {
		t.Errorf("response int64s length mismatch: %d != %d", len(resp.FInt64s), len(req.FInt64s))
	}
	for i, n := range resp.FInt64s {
		if n != req.FInt64s[i] {
			t.Errorf("response int64s mismatch at index %d: %d != %d", i, n, req.FInt64s[i])
		}
	}
	if !bytes.Equal(resp.FBytes, req.FBytes) {
		t.Errorf("response bytes mismatch: %s != %s", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("response bytsess length mismatch: %d != %d", len(resp.FBytess), len(req.FBytess))
	}
	for i, b := range resp.FBytess {
		if !bytes.Equal(b, req.FBytess[i]) {
			t.Errorf("response bytsess mismatch at index %d: %s != %s", i, b, req.FBytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("response float mismatch: %f != %f", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("response floats length mismatch: %d != %d", len(resp.FFloats), len(req.FFloats))
	}
	for i, f := range resp.FFloats {
		if f != req.FFloats[i] {
			t.Errorf("response floats mismatch at index %d: %f != %f", i, f, req.FFloats[i])
		}
	}
}
