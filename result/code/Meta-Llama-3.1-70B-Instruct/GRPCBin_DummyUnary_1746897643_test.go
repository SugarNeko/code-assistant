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

func TestDummyUnary(t *testing.T) {
	ctx := context.Background()

	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Create a request message
	req := &pb.DummyMessage{
		FString: "Hello, world!",
		FStrings: []string{
			"string 1",
			"string 2",
		},
		FInt32: 123,
		FInt32s: []int32{
			1, 2, 3,
		},
		Enum: pb.DummyMessage_ENUM_1,
		FEnums: []pb.DummyMessage_Enum{
			pb.DummyMessage_ENUM_0,
			pb.DummyMessage_ENUM_2,
		},
		FSub: &pb.DummyMessage_Sub{
			FString: "sub string",
		},
		FSubs: []*pb.DummyMessage_Sub{
			{FString: "sub string 1"},
			{FString: "sub string 2"},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64: 1234567890,
		FInt64s: []int64{
			1, 2, 3,
		},
		FBytes: []byte("bytes"),
		FBytess: [][]byte{
			[]byte("bytes 1"),
			[]byte("bytes 2"),
		},
		FFloat: 3.14,
		FFloats: []float32{
			1.1,
			2.2,
		},
	}

	// Call the gRPC method
	startTime := time.Now()
	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() != codes.OK {
				t.Errorf("grpc error: %v", st)
			}
		} else {
			t.Errorf("grpc error: %v", err)
		}
	}
	t.Logf("grpc response time: %v", time.Since(startTime))

	// Validate the response message
	if resp == nil {
		t.Errorf("response message is nil")
	}
	if resp.FString != req.FString {
		t.Errorf("FString does not match: expected %q, got %q", req.FString, resp.FString)
	}
	if !cmpSliceString(resp.FStrings, req.FStrings) {
		t.Errorf("FStrings does not match: expected %+v, got %+v", req.FStrings, resp.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 does not match: expected %d, got %d", req.FInt32, resp.FInt32)
	}
	if !cmpSliceInt32(resp.FInt32s, req.FInt32s) {
		t.Errorf("FInt32s does not match: expected %+v, got %+v", req.FInt32s, resp.FInt32s)
	}
	if resp.Enum != req.Enum {
		t.Errorf("Enum does not match: expected %d, got %d", req.Enum, resp.Enum)
	}
	if !cmpSliceEnum(resp.FEnums, req.FEnums) {
		t.Errorf("FEnums does not match: expected %+v, got %+v", req.FEnums, resp.FEnums)
	}
	if resp.FSub == nil {
		t.Errorf("FSub is nil")
	} else if resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub FString does not match: expected %q, got %q", req.FSub.FString, resp.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length does not match: expected %d, got %d", len(req.FSubs), len(resp.FSubs))
	} else if !cmpSliceSub(resp.FSubs, req.FSubs) {
		t.Errorf("FSubs does not match: expected %+v, got %+v", req.FSubs, resp.FSubs)
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool does not match: expected %t, got %t", req.FBool, resp.FBool)
	}
	if !cmpSliceBool(resp.FBools, req.FBools) {
		t.Errorf("FBools does not match: expected %+v, got %+v", req.FBools, resp.FBools)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 does not match: expected %d, got %d", req.FInt64, resp.FInt64)
	}
	if !cmpSliceInt64(resp.FInt64s, req.FInt64s) {
		t.Errorf("FInt64s does not match: expected %+v, got %+v", req.FInt64s, resp.FInt64s)
	}
	if !cmpSliceByte(resp.FBytes, req.FBytes) {
		t.Errorf("FBytes does not match: expected %+v, got %+v", req.FBytes, resp.FBytes)
	}
	if !cmpSliceByte(resp.FBytess, req.FBytess) {
		t.Errorf("FBytess does not match: expected %+v, got %+v", req.FBytess, resp.FBytess)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat does not match: expected %f, got %f", req.FFloat, resp.FFloat)
	}
	if !cmpSliceFloat(resp.FFloats, req.FFloats) {
		t.Errorf("FFloats does not match: expected %+v, got %+v", req.FFloats, resp.FFloats)
	}
}

func cmpSliceString(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cmpSliceInt32(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cmpSliceEnum(a, b []pb.DummyMessage_Enum) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cmpSliceSub(a, b []*pb.DummyMessage_Sub) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i].FString != b[i].FString {
			return false
		}
	}
	return true
}

func cmpSliceBool(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cmpSliceInt64(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cmpSliceByte(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func cmpSliceFloat(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
