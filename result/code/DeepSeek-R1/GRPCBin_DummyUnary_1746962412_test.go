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
		FInt32s:    []int32{1, 2, 3},
		FEnum:      grpcbin.DummyMessage_ENUM_1,
		FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:       &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:      true,
		FBools:     []bool{true, false, true},
		FInt64:     456,
		FInt64s:    []int64{4, 5, 6},
		FBytes:     []byte{0x01, 0x02},
		FBytess:    [][]byte{{0x03}, {0x04}},
		FFloat:     3.14,
		FFloats:    []float32{1.1, 2.2},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Validate response matches request
	if res.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", res.FString, req.FString)
	}

	if len(res.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(res.FStrings), len(req.FStrings))
	} else {
		for i := range res.FStrings {
			if res.FStrings[i] != req.FStrings[i] {
				t.Errorf("FStrings[%d] mismatch: got %q, want %q", i, res.FStrings[i], req.FStrings[i])
			}
		}
	}

	if res.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", res.FInt32, req.FInt32)
	}

	if len(res.FInt32s) != len(req.FInt32s) {
		t.Errorf("FInt32s length mismatch: got %d, want %d", len(res.FInt32s), len(req.FInt32s))
	} else {
		for i := range res.FInt32s {
			if res.FInt32s[i] != req.FInt32s[i] {
				t.Errorf("FInt32s[%d] mismatch: got %d, want %d", i, res.FInt32s[i], req.FInt32s[i])
			}
		}
	}

	if res.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", res.FEnum, req.FEnum)
	}

	if len(res.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(res.FEnums), len(req.FEnums))
	} else {
		for i := range res.FEnums {
			if res.FEnums[i] != req.FEnums[i] {
				t.Errorf("FEnums[%d] mismatch: got %v, want %v", i, res.FEnums[i], req.FEnums[i])
			}
		}
	}

	if res.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %q, want %q", res.FSub.FString, req.FSub.FString)
	}

	if len(res.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(res.FSubs), len(req.FSubs))
	} else {
		for i := range res.FSubs {
			if res.FSubs[i].FString != req.FSubs[i].FString {
				t.Errorf("FSubs[%d].FString mismatch: got %q, want %q", i, res.FSubs[i].FString, req.FSubs[i].FString)
			}
		}
	}

	if res.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", res.FBool, req.FBool)
	}

	if len(res.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(res.FBools), len(req.FBools))
	} else {
		for i := range res.FBools {
			if res.FBools[i] != req.FBools[i] {
				t.Errorf("FBools[%d] mismatch: got %v, want %v", i, res.FBools[i], req.FBools[i])
			}
		}
	}

	if res.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", res.FInt64, req.FInt64)
	}

	if len(res.FInt64s) != len(req.FInt64s) {
		t.Errorf("FInt64s length mismatch: got %d, want %d", len(res.FInt64s), len(req.FInt64s))
	} else {
		for i := range res.FInt64s {
			if res.FInt64s[i] != req.FInt64s[i] {
				t.Errorf("FInt64s[%d] mismatch: got %d, want %d", i, res.FInt64s[i], req.FInt64s[i])
			}
		}
	}

	if !bytes.Equal(res.FBytes, req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", res.FBytes, req.FBytes)
	}

	if len(res.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(res.FBytess), len(req.FBytess))
	} else {
		for i := range res.FBytess {
			if !bytes.Equal(res.FBytess[i], req.FBytess[i]) {
				t.Errorf("FBytess[%d] mismatch: got %v, want %v", i, res.FBytess[i], req.FBytess[i])
			}
		}
	}

	if res.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %f, want %f", res.FFloat, req.FFloat)
	}

	if len(res.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(res.FFloats), len(req.FFloats))
	} else {
		for i := range res.FFloats {
			if res.FFloats[i] != req.FFloats[i] {
				t.Errorf("FFloats[%d] mismatch: got %f, want %f", i, res.FFloats[i], req.FFloats[i])
			}
		}
	}
}
