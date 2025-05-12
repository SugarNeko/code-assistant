package grpcbin_test

import (
	"context"
	"testing"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:    "test",
		FStrings:   []string{"a", "b"},
		FInt32:     123,
		FInt32S:    []int32{1, 2},
		FEnum:      grpcbin.DummyMessage_ENUM_1,
		FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:       &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:      true,
		FBools:     []bool{true, false},
		FInt64:     456,
		FInt64S:    []int64{3, 4},
		FBytes:     []byte{0x01, 0x02},
		FBytess:    [][]byte{{0x03}, {0x04}},
		FFloat:     1.23,
		FFloats:    []float32{5.6, 7.8},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("DummyUnary RPC failed: %v", err)
	}

	// Response validation
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %v, want %v", resp.FString, req.FString)
	}

	comparePrimitiveSlices(t, "FStrings", resp.FStrings, req.FStrings)
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
	}
	compareInt32Slices(t, "FInt32s", resp.FInt32S, req.FInt32S)
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}
	compareEnumSlices(t, "FEnums", resp.FEnums, req.FEnums)
	compareSubMessages(t, "FSub", resp.FSub, req.FSub)
	compareSubSlices(t, "FSubs", resp.FSubs, req.FSubs)
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}
	compareBoolSlices(t, "FBools", resp.FBools, req.FBools)
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, req.FInt64)
	}
	compareInt64Slices(t, "FInt64s", resp.FInt64S, req.FInt64S)
	compareBytes(t, "FBytes", resp.FBytes, req.FBytes)
	compareByteSlices(t, "FBytess", resp.FBytess, req.FBytess)
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, req.FFloat)
	}
	compareFloatSlices(t, "FFloats", resp.FFloats, req.FFloats)
}

func comparePrimitiveSlices(t *testing.T, name string, got, want []string) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d] mismatch: got %v, want %v", name, i, got[i], want[i])
		}
	}
}

func compareInt32Slices(t *testing.T, name string, got, want []int32) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d] mismatch: got %v, want %v", name, i, got[i], want[i])
		}
	}
}

func compareEnumSlices(t *testing.T, name string, got, want []grpcbin.DummyMessage_Enum) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d] mismatch: got %v, want %v", name, i, got[i], want[i])
		}
	}
}

func compareSubMessages(t *testing.T, name string, got, want *grpcbin.DummyMessage_Sub) {
	if got.FString != want.FString {
		t.Errorf("%s.FString mismatch: got %v, want %v", name, got.FString, want.FString)
	}
}

func compareSubSlices(t *testing.T, name string, got, want []*grpcbin.DummyMessage_Sub) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		compareSubMessages(t, name, got[i], want[i])
	}
}

func compareBoolSlices(t *testing.T, name string, got, want []bool) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d] mismatch: got %v, want %v", name, i, got[i], want[i])
		}
	}
}

func compareInt64Slices(t *testing.T, name string, got, want []int64) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d] mismatch: got %v, want %v", name, i, got[i], want[i])
		}
	}
}

func compareBytes(t *testing.T, name string, got, want []byte) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d] mismatch: got %v, want %v", name, i, got[i], want[i])
		}
	}
}

func compareByteSlices(t *testing.T, name string, got, want [][]byte) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		compareBytes(t, name, got[i], want[i])
	}
}

func compareFloatSlices(t *testing.T, name string, got, want []float32) {
	if len(got) != len(want) {
		t.Errorf("%s length mismatch: got %d, want %d", name, len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d] mismatch: got %v, want %v", name, i, got[i], want[i])
		}
	}
}
