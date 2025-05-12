package grpcbin

import (
	"context"
	"google.golang.org/grpc"
	"testing"
)

const (
	grpcServerAddr string = "grpcb.in:9000"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	req := &DummyMessage{
		FString: "test",
		FStrings: []string{
			"test1",
			"test2",
		},
		FInt32: 123,
		FInt32s: []int32{
			1,
			2,
			3,
		},
		FEenum: ENUM_0,
		FEnums: []Enum{
			ENUM_1,
			ENUM_2,
		},
		FSub: &DummyMessage_Sub{
			FString: "sub-test",
		},
		FSubs: []*DummyMessage_Sub{
			{
				FString: "sub-test1",
			},
			{
				FString: "sub-test2",
			},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64:  123456,
		FInt64s: []int64{
			111,
			222,
			333,
		},
		FBytes: []byte("hoge"),
		FBytess: [][]byte{
			[]byte("hogehoge"),
			[]byte("fugafuga"),
		},
		FFloat:  1.23,
		FFloats: []float32{
			1.23,
			4.56,
		},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("DummyUnary error: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString unmatch, want %v, got %v", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length unmatch, want %v, got %v", len(req.FStrings), len(resp.FStrings))
	}
	for i, s := range resp.FStrings {
		if s != req.FStrings[i] {
			t.Errorf("FStrings[%v] unmatch, want %v, got %v", i, req.FStrings[i], s)
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 unmatch, want %v, got %v", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("FInt32s length unmatch, want %v, got %v", len(req.FInt32s), len(resp.FInt32s))
	}
	for i, n := range resp.FInt32s {
		if n != req.FInt32s[i] {
			t.Errorf("FInt32s[%v] unmatch, want %v, got %v", i, req.FInt32s[i], n)
		}
	}
	if resp.FEenum != req.FEenum {
		t.Errorf("FEenum unmatch, want %v, got %v", req.FEenum, resp.FEenum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length unmatch, want %v, got %v", len(req.FEnums), len(resp.FEnums))
	}
	for i, e := range resp.FEnums {
		if e != req.FEnums[i] {
			t.Errorf("FEnums[%v] unmatch, want %v, got %v", i, req.FEnums[i], e)
		}
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub FString unmatch, want %v, got %v", req.FSub.FString, resp.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length unmatch, want %v, got %v", len(req.FSubs), len(resp.FSubs))
	}
	for i, sub := range resp.FSubs {
		if sub.FString != req.FSubs[i].FString {
			t.Errorf("FSubs[%v] FString unmatch, want %v, got %v", i, req.FSubs[i].FString, sub.FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool unmatch, want %v, got %v", req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length unmatch, want %v, got %v", len(req.FBools), len(resp.FBools))
	}
	for i, b := range resp.FBools {
		if b != req.FBools[i] {
			t.Errorf("FBools[%v] unmatch, want %v, got %v", i, req.FBools[i], b)
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 unmatch, want %v, got %v", req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64s) != len(req.FInt64s) {
		t.Errorf("FInt64s length unmatch, want %v, got %v", len(req.FInt64s), len(resp.FInt64s))
	}
	for i, n := range resp.FInt64s {
		if n != req.FInt64s[i] {
			t.Errorf("FInt64s[%v] unmatch, want %v, got %v", i, req.FInt64s[i], n)
		}
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes unmatch, want %v, got %v", req.FBytes, resp.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length unmatch, want %v, got %v", len(req.FBytess), len(resp.FBytess))
	}
	for i, bs := range resp.FBytess {
		if string(bs) != string(req.FBytess[i]) {
			t.Errorf("FBytess[%v] unmatch, want %v, got %v", i, req.FBytess[i], bs)
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat unmatch, want %v, got %v", req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length unmatch, want %v, got %v", len(req.FFloats), len(resp.FFloats))
	}
	for i, f := range resp.FFloats {
		if f != req.FFloats[i] {
			t.Errorf("FFloats[%v] unmatch, want %v, got %v", i, req.FFloats[i], f)
		}
	}
}
