package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "code-assistant/proto/grpcbin"
)

const (
 address    = "grpcb.in:9000"
 modelName  = "GRPCBin"
)

func TestGRPCBinService_DummyUnary(t *testing.T) {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        t.Fatal(err)
    }
    defer conn.Close()

    client := pb.NewGRPCBinClient(conn)

    // Positive testing: Constructing typical requests
    req := &pb.DummyMessage{
        FString: "test",
        FStrings: []string{"test1", "test2"},
        FInt32: 1,
        FInt32s: []int32{1, 2},
        FEnum: pb.Enum_ENUM_0,
        FEnums: []pb.Enum{pb.Enum_ENUM_0, pb.Enum_ENUM_1},
        FSub: &pb.DummyMessage_Sub{
            FString: "test",
        },
        FSubs: []*pb.DummyMessage_Sub{
            {FString: "test1"},
            {FString: "test2"},
        },
        FBool: true,
        FBools: []bool{true, false},
        FInt64: 1,
        FInt64s: []int64{1, 2},
        FBytes: []byte("test"),
        FBytess: [][]byte{[]byte("test1"), []byte("test2")},
        FFloat: 1.0,
        FFloats: []float32{1.0, 2.0},
    }

    resp, err := client.DummyUnary(context.Background(), req)
    if err != nil {
        t.Fatal(err)
    }

    // Client response validation
    if resp.FString != req.FString {
        t.Errorf("resp.FString want %q, got %q", req.FString, resp.FString)
    }
    if !equalSliceString(resp.FStrings, req.FStrings) {
        t.Errorf("resp.FStrings want %v, got %v", req.FStrings, resp.FStrings)
    }
    if resp.FInt32 != req.FInt32 {
        t.Errorf("resp.FInt32 want %d, got %d", req.FInt32, resp.FInt32)
    }
    if !equalSliceInt32(resp.FInt32s, req.FInt32s) {
        t.Errorf("resp.FInt32s want %v, got %v", req.FInt32s, resp.FInt32s)
    }
    if resp.FEnum != req.FEnum {
        t.Errorf("resp.FEnum want %v, got %v", req.FEnum, resp.FEnum)
    }
    if !equalSliceEnum(resp.FEnums, req.FEnums) {
        t.Errorf("resp.FEnums want %v, got %v", req.FEnums, resp.FEnums)
    }
    if resp.FSub.FString != req.FSub.FString {
        t.Errorf("resp.FSub.FString want %q, got %q", req.FSub.FString, resp.FSub.FString)
    }
    if !equalSliceSub(resp.FSubs, req.FSubs) {
        t.Errorf("resp.FSubs want %v, got %v", req.FSubs, resp.FSubs)
    }
    if resp.FBool != req.FBool {
        t.Errorf("resp.FBool want %v, got %v", req.FBool, resp.FBool)
    }
    if !equalSliceBool(resp.FBools, req.FBools) {
        t.Errorf("resp.FBools want %v, got %v", req.FBools, resp.FBools)
    }
    if resp.FInt64 != req.FInt64 {
        t.Errorf("resp.FInt64 want %d, got %d", req.FInt64, resp.FInt64)
    }
    if !equalSliceInt64(resp.FInt64s, req.FInt64s) {
        t.Errorf("resp.FInt64s want %v, got %v", req.FInt64s, resp.FInt64s)
    }
    if !equalSliceByte(resp.FBytes, req.FBytes) {
        t.Errorf("resp.FBytes want %v, got %v", req.FBytes, resp.FBytes)
    }
    if !equalSliceByte(resp.FBytess, req.FBytess) {
        t.Errorf("resp.FBytess want %v, got %v", req.FBytess, resp.FBytess)
    }
    if resp.FFloat != req.FFloat {
        t.Errorf("resp.FFloat want %f, got %f", req.FFloat, resp.FFloat)
    }
    if !equalSliceFloat(resp.FFloats, req.FFloats) {
        t.Errorf("resp.FFloats want %v, got %v", req.FFloats, resp.FFloats)
    }

    // Server response validation
    if status.Code(err) != codes.OK {
        t.Errorf("status code want %v, got %v", codes.OK, status.Code(err))
    }
}

// equalSliceString returns true if two slices of string are equal.
func equalSliceString(a, b []string) bool {
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

// equalSliceInt32 returns true if two slices of int32 are equal.
func equalSliceInt32(a, b []int32) bool {
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

// equalSliceEnum returns true if two slices of Enum are equal.
func equalSliceEnum(a, b []pb.Enum) bool {
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

// equalSliceSub returns true if two slices of Sub are equal.
func equalSliceSub(a, b []*pb.DummyMessage_Sub) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if !equalSub(a[i], b[i]) {
            return false
        }
    }
    return true
}

// equalSub returns true if two Sub are equal.
func equalSub(a, b *pb.DummyMessage_Sub) bool {
    if a.FString != b.FString {
        return false
    }
    return true
}

// equalSliceBool returns true if two slices of bool are equal.
func equalSliceBool(a, b []bool) bool {
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

// equalSliceInt64 returns true if two slices of int64 are equal.
func equalSliceInt64(a, b []int64) bool {
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

// equalSliceByte returns true if two slices of byte are equal.
func equalSliceByte(a, b []byte) bool {
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

// equalSliceByte returns true if two slices of byte are equal.
func equalSliceByte(a, b [][]byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if !equalSliceByte(a[i], b[i]) {
            return false
        }
    }
    return true
}

// equalSliceFloat returns true if two slices of float32 are equal.
func equalSliceFloat(a, b []float32) bool {
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
