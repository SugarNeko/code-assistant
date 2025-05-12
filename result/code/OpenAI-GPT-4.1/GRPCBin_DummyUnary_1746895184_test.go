package grpcbin_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test string",
		FStrings: []string{"str1", "str2"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub1"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub2"}, {FString: "sub3"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   1000000000,
		FInt64S:  []int64{100, 200},
		FBytes:   []byte("hello"),
		FBytess:  [][]byte{[]byte("a"), []byte("b")},
		FFloat:   1.23,
		FFloats:  []float32{4.56, 7.89},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	// Validate each field in the response equals the request
	if resp.FString != req.FString {
		t.Errorf("FString: got %q, want %q", resp.FString, req.FString)
	}
	if !equalStringSlice(resp.FStrings, req.FStrings) {
		t.Errorf("FStrings: got %v, want %v", resp.FStrings, req.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32: got %v, want %v", resp.FInt32, req.FInt32)
	}
	if !equalInt32Slice(resp.FInt32S, req.FInt32S) {
		t.Errorf("FInt32S: got %v, want %v", resp.FInt32S, req.FInt32S)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if !equalEnumSlice(resp.FEnums, req.FEnums) {
		t.Errorf("FEnums: got %v, want %v", resp.FEnums, req.FEnums)
	}
	if (resp.FSub == nil) != (req.FSub == nil) || (resp.FSub != nil && resp.FSub.FString != req.FSub.FString) {
		t.Errorf("FSub: got %v, want %v", resp.FSub, req.FSub)
	}
	if !equalSubSlice(resp.FSubs, req.FSubs) {
		t.Errorf("FSubs: got %v, want %v", resp.FSubs, req.FSubs)
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool: got %v, want %v", resp.FBool, req.FBool)
	}
	if !equalBoolSlice(resp.FBools, req.FBools) {
		t.Errorf("FBools: got %v, want %v", resp.FBools, req.FBools)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64: got %v, want %v", resp.FInt64, req.FInt64)
	}
	if !equalInt64Slice(resp.FInt64S, req.FInt64S) {
		t.Errorf("FInt64S: got %v, want %v", resp.FInt64S, req.FInt64S)
	}
	if !bytes.Equal(resp.FBytes, req.FBytes) {
		t.Errorf("FBytes: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if !equalBytesSlice(resp.FBytess, req.FBytess) {
		t.Errorf("FBytess: got %v, want %v", resp.FBytess, req.FBytess)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat: got %v, want %v", resp.FFloat, req.FFloat)
	}
	if !equalFloat32Slice(resp.FFloats, req.FFloats) {
		t.Errorf("FFloats: got %v, want %v", resp.FFloats, req.FFloats)
	}
}

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalInt32Slice(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalEnumSlice(a, b []grpcbin.DummyMessage_Enum) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalSubSlice(a, b []*grpcbin.DummyMessage_Sub) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if (a[i] == nil) != (b[i] == nil) {
			return false
		}
		if a[i] != nil && a[i].FString != b[i].FString {
			return false
		}
	}
	return true
}

func equalBoolSlice(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalInt64Slice(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalBytesSlice(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !bytes.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalFloat32Slice(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
