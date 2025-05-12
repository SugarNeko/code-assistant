package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin/grpcbin"
)

const (
	address = "grpcb.in:9000"
)

func TestDummyUnary(t *testing.T) {
	cc, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()

	client := pb.NewGRPCBinClient(cc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dummyMessage := &pb.DummyMessage{
		FString:  "dummy_string",
		FStrings: []string{"dummy_string1", "dummy_string2"},
		FInt32:   123,
		FInt32s:  []int32{123, 456},
		FFields:  pb.DummyMessage.Enum.pbりの(Enum_1),
		FFields:  []pb.DummyMessageDTO.Fields{pb.DummyMessage.Fields(Fields_1)},
		FSub: &pb.DummyMessageSub{
			FString: "dummy_sub_string",
		},
		FSubs: []*pb.DummyMessageSub{
			{
				FString: "dummy_sub_string1",
			},
			{
				FString: "dummy_sub_string2",
			},
		},
		FBool:  true,
		FBools: []bool{true, false},
		FInt64: 123456,
		FInt64s: []int64{123456, 789012},
		FBytes:  []byte("dummy_bytes"),
		FBytes:  [][]byte{[]byte("dummy_bytes1"), []byte("dummy_bytes2")},
		FFloat:  123.456,
		FFloats: []float32{123.456, 789.012},
	}
	resp, err := client.DummyUnary(ctx, dummyMessage)
	if err != nil {
		t.Errorf("could not greet: %v", err)
	}
	if resp == nil {
		t.Errorf("empty response")
	}
	if resp.FString != dummyMessage.FString {
		t.Errorf("want %s, got %s", dummyMessage.FString, resp.FString)
	}
	if !EqualStringSlices(resp.FStrings, dummyMessage.FStrings) {
		t.Errorf("want %v, got %v", dummyMessage.FStrings, resp.FStrings)
	}
	if resp.FInt32 != dummyMessage.FInt32 {
		t.Errorf("want %d, got %d", dummyMessage.FInt32, resp.FInt32)
	}
	if !EqualInt32Slices(resp.FInt32s, dummyMessage.FInt32s) {
		t.Errorf("want %v, got %v", dummyMessage.FInt32s, resp.FInt32s)
	}
	if resp.FEnum != dummyMessage.FEnum {
		t.Errorf("want %s, got %s", dummyMessage.FEnum, resp.FEnum)
	}
	if !EqualEnumSlices(resp.FEnums, dummyMessage.FEnums) {
		t.Errorf("want %v, got %v", dummyMessage.FEnums, resp.FEnums)
	}
	if !EqualSubMessages(resp.FSub, dummyMessage.FSub) {
		t.Errorf("want %v, got %v", dummyMessage.FSub, resp.FSub)
	}
	if !EqualSubMessagesSlices(resp.FSubs, dummyMessage.FSubs) {
		t.Errorf("want %v, got %v", dummyMessage.FSubs, resp.FSubs)
	}
	if resp.FBool != dummyMessage.FBool {
		t.Errorf("want %t, got %t", dummyMessage.FBool, resp.FBool)
	}
	if !EqualBoolSlices(resp.FBools, dummyMessage.FBools) {
		t.Errorf("want %v, got %v", dummyMessage.FBools, resp.FBools)
	}
	if resp.FInt64 != dummyMessage.FInt64 {
		t.Errorf("want %d, got %d", dummyMessage.FInt64, resp.FInt64)
	}
	if !EqualInt64Slices(resp.FInt64s, dummyMessage.FInt64s) {
		t.Errorf("want %v, got %v", dummyMessage.FInt64s, resp.FInt64s)
	}
	if !EqualBytesSlices(resp.FBytes, dummyMessage.FBytes) {
		t.Errorf("want %v, got %v", dummyMessage.FBytes, resp.FBytes)
	}
	if resp.FFloat != dummyMessage.FFloat {
		t.Errorf("want %f, got %f", dummyMessage.FFloat, resp.FFloat)
	}
	if !EqualFloat32Slices(resp.FFloats, dummyMessage.FFloats) {
		t.Errorf("want %v, got %v", dummyMessage.FFloats, resp.FFloats)
	}
}

func EqualStringSlices(a, b []string) bool {
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

func EqualInt32Slices(a, b []int32) bool {
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

func EqualEnumSlices(a, b []pb.DummyMessageDTO.Fields) bool {
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

func EqualSubMessages(a, b *pb.DummyMessageSub) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.FString == b.FString
}

func EqualSubMessagesSlices(a, b []*pb.DummyMessageSub) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !EqualSubMessages(v, b[i]) {
			return false
		}
	}
	return true
}

func EqualBoolSlices(a, b []bool) bool {
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

func EqualInt64Slices(a, b []int64) bool {
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

func EqualBytesSlices(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !EqualBytes(v, b[i]) {
			return false
		}
	}
	return true
}

func EqualBytes(a, b []byte) bool {
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

func EqualFloat32Slices(a, b []float32) bool {
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
