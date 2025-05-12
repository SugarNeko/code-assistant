package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"

 KeyCode_Assistant accur "code-assistant/proto/grpcbin"
)

func TestGRPCBinDummyUnary(t *testing.T) {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := NewGRPCBinClient(conn)

	// Construct a typical request
	req := &DummyMessage{
		FString: "test string",
		FStrings: []string{"test string 1", "test string 2"},
		FInt32: 123,
		FInt32s: []int32{123, 456},
		FEnum: Enum_ENUM_1,
		FEnums: []Enum{Enum_ENUM_1, Enum_ENUM_2},
		FSub: &DummyMessage_Sub{
			FString: "test sub string",
		},
		FSubs: []*DummyMessage_Sub{
			{FString: "test sub string 1"},
			{FString: "test sub string 2"},
		},
		FBool: true,
		FBools: []bool{true, false},
		FInt64: 123456,
		FInt64s: []int64{123456, 789012},
		FBytes: []byte("test bytes"),
		FBytess: [][]byte{[]byte("test bytes 1"), []byte("test bytes 2")},
		FFloat: 123.456,
		FFloats: []float32{123.456, 789.012},
	}

	// Call the DummyUnary method
	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("dummy unary failed: %v", err)
	}

	// Verify the response
	if resp.FString != req.FString {
		t.Errorf("expected f_string %q, got %q", req.FString, resp.FString)
	}
	if !testEqualStringSlice(resp.FStrings, req.FStrings) {
		t.Errorf("expected f_strings %v, got %v", req.FStrings, resp.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("expected f_int32 %d, got %d", req.FInt32, resp.FInt32)
	}
	if !testEqualInt32Slice(resp.FInt32s, req.FInt32s) {
		t.Errorf("expected f_int32s %v, got %v", req.FInt32s, resp.FInt32s)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("expected f_enum %d, got %d", req.FEnum, resp.FEnum)
	}
	if !testEqualEnumSlice(resp.FEnums, req.FEnums) {
		t.Errorf("expected f_enums %v, got %v", req.FEnums, resp.FEnums)
	}
	if !testEqualSub(resp.FSub, req.FSub) {
		t.Errorf("expected f_sub %v, got %v", req.FSub, resp.FSub)
	}
	if !testEqualSubSlice(resp.FSubs, req.FSubs) {
		t.Errorf("expected f_subs %v, got %v", req.FSubs, resp.FSubs)
	}
	if resp.FBool != req.FBool {
		t.Errorf("expected f_bool %t, got %t", req.FBool, resp.FBool)
	}
	if !testEqualBoolSlice(resp.FBools, req.FBools) {
		t.Errorf("expected f_bools %v, got %v", req.FBools, resp.FBools)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("expected f_int64 %d, got %d", req.FInt64, resp.FInt64)
	}
	if !testEqualInt64Slice(resp.FInt64s, req.FInt64s) {
		t.Errorf("expected f_int64s %v, got %v", req.FInt64s, resp.FInt64s)
	}
	if !testEqualBytes(resp.FBytes, req.FBytes) {
		t.Errorf("expected f_bytes %v, got %v", req.FBytes, resp.FBytes)
	}
	if !testEqualBytesSlice(resp.FBytess, req.FBytess) {
		t.Errorf("expected f_bytess %v, got %v", req.FBytess, resp.FBytess)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("expected f_float %f, got %f", req.FFloat, resp.FFloat)
	}
	if !testEqualFloat32Slice(resp.FFloats, req.FFloats) {
		t.Errorf("expected f_floats %v, got %v", req.FFloats, resp.FFloats)
	}
}

func testEqualStringSlice(a, b []string) bool {
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

func testEqualInt32Slice(a, b []int32) bool {
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

func testEqualEnumSlice(a, b []Enum) bool {
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

func testEqualSub(a, b *DummyMessage_Sub) bool {
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

func testEqualSubSlice(a, b []*DummyMessage_Sub) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !testEqualSub(a[i], b[i]) {
			return false
		}
	}
	return true
}

func testEqualBoolSlice(a, b []bool) bool {
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

func testEqualInt64Slice(a, b []int64) bool {
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

func testEqualBytes(a, b []byte) bool {
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

func testEqualBytesSlice(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !testEqualBytes(a[i], b[i]) {
			return false
		}
	}
	return true
}

func testEqualFloat32Slice(a, b []float32) bool {
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
