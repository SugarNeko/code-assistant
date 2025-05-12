package grpcbin

import (
	"context"
	"fmt"
	"log"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Constructing typical requests that fully comply with the interface specification
	req := &pb.DummyMessage{
		FString: "Hello, World!",
		FStrings: []string{
			"Hello",
			"World",
		},
		FInt32: 42,
		FInt32s: []int32{
			1,
			2,
			3,
		},
		FFieldMask: &pb.DummyMessage_Sub{
			FString: "Embedded Message",
		},
		FEnum: pb.DummyMessage_ENUM_1,
		FEnums: []pb.DummyMessage_Enum{
			pb.DummyMessage_ENUM_0,
			pb.DummyMessage_ENUM_1,
			pb.DummyMessage_ENUM_2,
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
			true,
		},
		FInt64: 64,
		FInt64s: []int64{
			64,
			128,
			256,
		},
		FBytes: []byte{'H', 'e', 'l', 'l', 'o'},
		FBytess: [][]byte{
			[]byte{'H', 'e', 'l', 'l', 'o'},
			[]byte{'W', 'o', 'r', 'l', 'd'},
		},
		FFloat: 3.14,
		FFloats: []float32{
			3.14,
			2.71,
			1.61,
		},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("error calling DummyUnary: %v", err)
	}

	// Client response validation
	if resp == nil {
		t.Errorf("response is nil")
	}
	if resp.FString != req.FString {
		t.Errorf("response string mismatch: want %q, got %q", req.FString, resp.FString)
	}
	if !compareSlices(req.FStrings, resp.FStrings) {
		t.Errorf("response string slices mismatch: want %v, got %v", req.FStrings, resp.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("response int32 mismatch: want %d, got %d", req.FInt32, resp.FInt32)
	}
	if !compareSlices(req.FInt32s, resp.FInt32s) {
		t.Errorf("response int32 slices mismatch: want %v, got %v", req.FInt32s, resp.FInt32s)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("response enum mismatch: want %d, got %d", req.FEnum, resp.FEnum)
	}
	if !compareSlices(req.FEnums, resp.FEnums) {
		t.Errorf("response enum slices mismatch: want %v, got %v", req.FEnums, resp.FEnums)
	}
	if !compareMessages(req.FFieldMask, resp.FFieldMask) {
		t.Errorf("response embedded message mismatch: want %v, got %v", req.FFieldMask, resp.FFieldMask)
	}
	if resp.FBool != req.FBool {
		t.Errorf("response bool mismatch: want %v, got %v", req.FBool, resp.FBool)
	}
	if !compareSlices(req.FBools, resp.FBools) {
		t.Errorf("response bool slices mismatch: want %v, got %v", req.FBools, resp.FBools)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("response int64 mismatch: want %d, got %d", req.FInt64, resp.FInt64)
	}
	if !compareSlices(req.FInt64s, resp.FInt64s) {
		t.Errorf("response int64 slices mismatch: want %v, got %v", req.FInt64s, resp.FInt64s)
	}
	if !compareByteSlices(req.FBytes, resp.FBytes) {
		t.Errorf("response bytes mismatch: want %v, got %v", req.FBytes, resp.FBytes)
	}
	if !compareByteSlices(req.FBytess, resp.FBytess) {
		t.Errorf("response bytes slices mismatch: want %v, got %v", req.FBytess, resp.FBytess)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("response float mismatch: want %f, got %f", req.FFloat, resp.FFloat)
	}
	if !compareSlices(req.FFloats, resp.FFloats) {
		t.Errorf("response float slices mismatch: want %v, got %v", req.FFloats, resp.FFloats)
	}
}

// Helper function to compare message fields
func compareMessages(a, b *pb.DummyMessage_Sub) bool {
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

// Helper function to compare byte slices
func compareByteSlices(a, b []byte) bool {
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

// Helper function to compare slices
func compareSlices[T any](a, b []T) bool {
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
