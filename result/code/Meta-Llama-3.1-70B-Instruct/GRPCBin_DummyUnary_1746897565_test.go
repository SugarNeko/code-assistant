package grpcbin

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	address = "grpcb.in:9000"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	req := &DummyMessage{
		FString: "hello",
		FStrings: []string{
			"world",
			"grpc",
		},
		FInt32: 42,
		FInt32s: []int32{
			24,
			42,
		},
		FEenum: Enum_ENUM_1,
		FEenms: []Enum{
			Enum_ENUM_1,
			Enum_ENUM_2,
		},
		FSub: &DummyMessage_Sub{
			FString: "sub",
		},
		FSubs: []*DummyMessage_Sub{
			{
				FString: "sub1",
			},
			{
				FString: "sub2",
			},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64:  42,
		FInt64s: []int64{
			24,
			42,
		},
		FBytes: []byte("bytes"),
		FBytess: [][]byte{
			[]byte("bytes1"),
			[]byte("bytes2"),
		},
		FFloat: 3.14,
		FFloats: []float32{
			1.23,
			4.56,
		},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("call Unary failed: %v", err)
	}

	if resp == nil {
		t.Errorf("response is nil")
	}

	if resp.FString != req.FString {
		t.Errorf("response FString = %s, want %s", resp.FString, req.FString)
	}
	if !cmpStringSlice(resp.FStrings, req.FStrings) {
		t.Errorf("response FStrings = %v, want %v", resp.FStrings, req.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("response FInt32 = %d, want %d", resp.FInt32, req.FInt32)
	}
	if !cmpInt32Slice(resp.FInt32s, req.FInt32s) {
		t.Errorf("response FInt32s = %v, want %v", resp.FInt32s, req.FInt32s)
	}
	if resp.FEenum != req.FEenum {
		t.Errorf("response FEenum = %v, want %v", resp.FEenum, req.FEenum)
	}
	if !cmpEnumSlice(resp.FEenms, req.FEenms) {
		t.Errorf("response FEenms = %v, want %v", resp.FEenms, req.FEenms)
	}
	if resp.FSub == nil {
		t.Errorf("response FSub is nil")
	}
	if !cmpDummyMessageSub(resp.FSub, req.FSub) {
		t.Errorf("response Fsub = %v, want %v", resp.FSub, req.FSub)
	}
	if !cmpDummyMessageSubSlice(resp.FSubs, req.FSubs) {
		t.Errorf("response FSubs = %v, want %v", resp.FSubs, req.FSubs)
	}
	if resp.FBool != req.FBool {
		t.Errorf("response FBool = %v, want %v", resp.FBool, req.FBool)
	}
	if !cmpBoolSlice(resp.FBools, req.FBools) {
		t.Errorf("response FBools = %v, want %v", resp.FBools, req.FBools)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("response FInt64 = %d, want %d", resp.FInt64, req.FInt64)
	}
	if !cmpInt64Slice(resp.FInt64s, req.FInt64s) {
		t.Errorf("response FInt64s = %v, want %v", resp.FInt64s, req.FInt64s)
	}
	if !cmpByteSlice(resp.FBytes, req.FBytes) {
		t.Errorf("response FBytes = %v, want %v", resp.FBytes, req.FBytes)
	}
	if !cmpByteArraySlice(resp.FBytess, req.FBytess) {
		t.Errorf("response FBytess = %v, want %v", resp.FBytess, req.FBytess)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("response FFloat = %f, want %f", resp.FFloat, req.FFloat)
	}
	if !cmpFloatSlice(resp.FFloats, req.FFloats) {
		t.Errorf("response FFloats = %v, want %v", resp.FFloats, req.FFloats)
	}
}

func cmpStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func cmpInt32Slice(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func cmpEnumSlice(a, b []Enum) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func cmpDummyMessageSub(a, b *DummyMessage_Sub) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.FString != b.FString {
		return false
	}
	return true
}

func cmpDummyMessageSubSlice(a, b []*DummyMessage_Sub) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !cmpDummyMessageSub(v, b[i]) {
			return false
		}
	}
	return true
}

func cmpBoolSlice(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func cmpInt64Slice(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func cmpByteSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func cmpByteArraySlice(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !cmpByteSlice(v, b[i]) {
			return false
		}
	}
	return true
}

func cmpFloatSlice(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
