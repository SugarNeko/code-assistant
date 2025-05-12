package grpcbin

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestGRPCBinDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewGRPCBinClient(conn)

	dummyMessage := &DummyMessage{
		FString: "hello",
		FStrings: []string{
			"hello",
			"world",
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
		FInt64: 1234567890,
		FInt64s: []int64{
			1234567890,
			9876543210,
		},
		FBytes: []byte{
			0x12, 0x34,
		},
		FByteses: [][]byte{
			{
				0x12, 0x34,
			},
			{
				0x56, 0x78,
			},
		},
		FFloat: 123.456,
		FFloats: []float32{
			123.456,
			789.012,
		},
	}

	resp, err := client.DummyUnary(context.Background(), dummyMessage)
	if err != nil {
		t.Fatalf("DummyUnary(_) = _, %v, want nil", err)
	}

	if resp.FString != dummyMessage.FString {
		t.Errorf("resp.FString = %s, want %s", resp.FString, dummyMessage.FString)
	}
	if len(resp.FStrings) != len(dummyMessage.FStrings) {
		t.Errorf("len(resp.FStrings) = %d, want %d", len(resp.FStrings), len(dummyMessage.FStrings))
	}
	for i, v := range resp.FStrings {
		if v != dummyMessage.FStrings[i] {
			t.Errorf("resp.FStrings[%d] = %s, want %s", i, v, dummyMessage.FStrings[i])
		}
	}
	if resp.FInt32 != dummyMessage.FInt32 {
		t.Errorf("resp.FInt32 = %d, want %d", resp.FInt32, dummyMessage.FInt32)
	}
	if len(resp.FInt32s) != len(dummyMessage.FInt32s) {
		t.Errorf("len(resp.FInt32s) = %d, want %d", len(resp.FInt32s), len(dummyMessage.FInt32s))
	}
	for i, v := range resp.FInt32s {
		if v != dummyMessage.FInt32s[i] {
			t.Errorf("resp.FInt32s[%d] = %d, want %d", i, v, dummyMessage.FInt32s[i])
		}
	}
	if resp.FEnum != dummyMessage.FEnum {
		t.Errorf("resp.FEnum = %s, want %s", resp.FEnum, dummyMessage.FEnum)
	}
	if len(resp.FEnums) != len(dummyMessage.FEnums) {
		t.Errorf("len(resp.FEnums) = %d, want %d", len(resp.FEnums), len(dummyMessage.FEnums))
	}
	for i, v := range resp.FEnums {
		if v != dummyMessage.FEnums[i] {
			t.Errorf("resp.FEnums[%d] = %s, want %s", i, v, dummyMessage.FEnums[i])
		}
	}
	if resp.FSub.FString != dummyMessage.FSub.FString {
		t.Errorf("resp.FSub.FString = %s, want %s", resp.FSub.FString, dummyMessage.FSub.FString)
	}
	if len(resp.FSubs) != len(dummyMessage.FSubs) {
		t.Errorf("len(resp.FSubs) = %d, want %d", len(resp.FSubs), len(dummyMessage.FSubs))
	}
	for i, v := range resp.FSubs {
		if v.FString != dummyMessage.FSubs[i].FString {
			t.Errorf("resp.FSubs[%d].FString = %s, want %s", i, v.FString, dummyMessage.FSubs[i].FString)
		}
	}
	if resp.FBool != dummyMessage.FBool {
		t.Errorf("resp.FBool = %v, want %v", resp.FBool, dummyMessage.FBool)
	}
	if len(resp.FBools) != len(dummyMessage.FBools) {
		t.Errorf("len(resp.FBools) = %d, want %d", len(resp.FBools), len(dummyMessage.FBools))
	}
	for i, v := range resp.FBools {
		if v != dummyMessage.FBools[i] {
			t.Errorf("resp.FBools[%d] = %v, want %v", i, v, dummyMessage.FBools[i])
		}
	}
	if resp.FInt64 != dummyMessage.FInt64 {
		t.Errorf("resp.FInt64 = %d, want %d", resp.FInt64, dummyMessage.FInt64)
	}
	if len(resp.FInt64s) != len(dummyMessage.FInt64s) {
		t.Errorf("len(resp.FInt64s) = %d, want %d", len(resp.FInt64s), len(dummyMessage.FInt64s))
	}
	for i, v := range resp.FInt64s {
		if v != dummyMessage.FInt64s[i] {
			t.Errorf("resp.FInt64s[%d] = %d, want %d", i, v, dummyMessage.FInt64s[i])
		}
	}
	if string(resp.FBytes) != string(dummyMessage.FBytes) {
		t.Errorf("resp.FBytes = %s, want %s", resp.FBytes, dummyMessage.FBytes)
	}
	if len(resp.FByteses) != len(dummyMessage.FByteses) {
		t.Errorf("len(resp.FByteses) = %d, want %d", len(resp.FByteses), len(dummyMessage.FByteses))
	}
	for i, v := range resp.FByteses {
		if string(v) != string(dummyMessage.FByteses[i]) {
			t.Errorf("resp.FByteses[%d] = %s, want %s", i, v, dummyMessage.FByteses[i])
		}
	}
	if resp.FFloat != dummyMessage.FFloat {
		t.Errorf("resp.FFloat = %f, want %f", resp.FFloat, dummyMessage.FFloat)
	}
	if len(resp.FFloats) != len(dummyMessage.FFloats) {
		t.Errorf("len(resp.FFloats) = %d, want %d", len(resp.FFloats), len(dummyMessage.FFloats))
	}
	for i, v := range resp.FFloats {
		if v != dummyMessage.FFloats[i] {
			t.Errorf("resp.FFloats[%d] = %f, want %f", i, v, dummyMessage.FFloats[i])
		}
	}
}
