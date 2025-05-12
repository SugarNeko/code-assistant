package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	req := &DummyMessage{
		FString: "Hello",
		FStrings: []string{
			"Hello",
			"World",
		},
		FInt32: 123,
		FInt32s: []int32{
			123,
			456,
		},
		FEnum: Enum_ENUM_1,
		FEnums: []Enum{
			Enum_ENUM_1,
			Enum_ENUM_2,
		},
		FSub: &DummyMessage_Sub{
			FString: "Hello",
		},
		FSubs: []*DummyMessage_Sub{
			{
				FString: "Hello",
			},
			{
				FString: "World",
			},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64: 123,
		FInt64s: []int64{
			123,
			456,
		},
		FBytes: []byte("Hello"),
		FBytess: [][]byte{
			[]byte("Hello"),
			[]byte("World"),
		},
		FFloat: 3.14,
		FFloats: []float32{
			3.14,
			2.71,
		},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString mismatch: want %s, got %s", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings mismatch: want %d, got %d", len(req.FStrings), len(resp.FStrings))
	}
	for i, s := range req.FStrings {
		if resp.FStrings[i] != s {
			t.Errorf("FStrings[%d] mismatch: want %s, got %s", i, s, resp.FStrings[i])
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: want %d, got %d", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("FInt32s mismatch: want %d, got %d", len(req.FInt32s), len(resp.FInt32s))
	}
	for i, s := range req.FInt32s {
		if resp.FInt32s[i] != s {
			t.Errorf("FInt32s[%d] mismatch: want %d, got %d", i, s, resp.FInt32s[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEunm mismatch: want %s, got %s", req.FEnum, resp.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums mismatch: want %d, got %d", len(req.FEnums), len(resp.FEnums))
	}
	for i, s := range req.FEnums {
		if resp.FEnums[i] != s {
			t.Errorf("FEnums[%d] mismatch: want %s, got %s", i, s, resp.FEnums[i])
		}
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: want %s, got %s", req.FSub.FString, resp.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs mismatch: want %d, got %d", len(req.FSubs), len(resp.FSubs))
	}
	for i, s := range req.FSubs {
		if resp.FSubs[i].FString != s.FString {
			t.Errorf("FSubs[%d].FString mismatch: want %s, got %s", i, s.FString, resp.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: want %t, got %t", req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools mismatch: want %d, got %d", len(req.FBools), len(resp.FBools))
	}
	for i, s := range req.FBools {
		if resp.FBools[i] != s {
			t.Errorf("FBools[%d] mismatch: want %t, got %t", i, s, resp.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: want %d, got %d", req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64s) != len(req.FInt64s) {
		t.Errorf("FInt64s mismatch: want %d, got %d", len(req.FInt64s), len(resp.FInt64s))
	}
	for i, s := range req.FInt64s {
		if resp.FInt64s[i] != s {
			t.Errorf("FInt64s[%d] mismatch: want %d, got %d", i, s, resp.FInt64s[i])
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: want %s, got %s", req.FBytes, resp.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess mismatch: want %d, got %d", len(req.FBytess), len(resp.FBytess))
	}
	for i, s := range req.FBytess {
		if string(resp.FBytess[i]) != string(s) {
			t.Errorf("FBytess[%d] mismatch: want %s, got %s", i, s, resp.FBytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: want %f, got %f", req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats mismatch: want %d, got %d", len(req.FFloats), len(resp.FFloats))
	}
	for i, s := range req.FFloats {
		if resp.FFloats[i] != s {
			t.Errorf("FFloats[%d] mismatch: want %f, got %f", i, s, resp.FFloats[i])
		}
	}
}
