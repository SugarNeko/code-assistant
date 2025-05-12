package grpcbin_test

import (
	"context"
	"testing"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"a", "b"},
		FInt32:   123,
		FInt32s:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "sub",
		},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{FString: "sub1"},
			{FString: "sub2"},
		},
		FBool:   true,
		FBools:  []bool{true, false},
		FInt64:  456,
		FInt64s: []int64{4, 5, 6},
		FBytes:  []byte("bytes"),
		FByteness: [][]byte{
			[]byte("b1"),
			[]byte("b2"),
		},
		FFloat:   1.23,
		FFloats:  []float32{3.14, 2.718},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary RPC failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
	}

	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
	} else {
		for i, v := range req.FStrings {
			if resp.FStrings[i] != v {
				t.Errorf("FStrings[%d] mismatch: got %q, want %q", i, resp.FStrings[i], v)
			}
		}
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
	}

	if len(resp.FInt32s) != len(req.FInt32s) {
		t.Errorf("FInt32s length mismatch: got %d, want %d", len(resp.FInt32s), len(req.FInt32s))
	} else {
		for i, v := range req.FInt32s {
			if resp.FInt32s[i] != v {
				t.Errorf("FInt32s[%d] mismatch: got %d, want %d", i, resp.FInt32s[i], v)
			}
		}
	}

	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}

	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(resp.FEnums), len(req.FEnums))
	} else {
		for i, v := range req.FEnums {
			if resp.FEnums[i] != v {
				t.Errorf("FEnums[%d] mismatch: got %v, want %v", i, resp.FEnums[i], v)
			}
		}
	}

	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %q, want %q", resp.FSub.FString, req.FSub.FString)
	}

	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
	} else {
		for i, v := range req.FSubs {
			if resp.FSubs[i].FString != v.FString {
				t.Errorf("FSubs[%d].FString mismatch: got %q, want %q", i, resp.FSubs[i].FString, v.FString)
			}
		}
	}

	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %t, want %t", resp.FBool, req.FBool)
	}

	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(resp.FBools), len(req.FBools))
	} else {
		for i, v := range req.FBools {
			if resp.FBools[i] != v {
				t.Errorf("FBools[%d] mismatch: got %t, want %t", i, resp.FBools[i], v)
			}
		}
	}

	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
	}

	if len(resp.FInt64s) != len(req.FInt64s) {
		t.Errorf("FInt64s length mismatch: got %d, want %d", len(resp.FInt64s), len(req.FInt64s))
	} else {
		for i, v := range req.FInt64s {
			if resp.FInt64s[i] != v {
				t.Errorf("FInt64s[%d] mismatch: got %d, want %d", i, resp.FInt64s[i], v)
			}
		}
	}

	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %q, want %q", resp.FBytes, req.FBytes)
	}

	if len(resp.FByteness) != len(req.FByteness) {
		t.Errorf("FByteness length mismatch: got %d, want %d", len(resp.FByteness), len(req.FByteness))
	} else {
		for i, v := range req.FByteness {
			if string(resp.FByteness[i]) != string(v) {
				t.Errorf("FByteness[%d] mismatch: got %q, want %q", i, resp.FByteness[i], v)
			}
		}
	}

	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %f, want %f", resp.FFloat, req.FFloat)
	}

	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(resp.FFloats), len(req.FFloats))
	} else {
		for i, v := range req.FFloats {
			if resp.FFloats[i] != v {
				t.Errorf("FFloats[%d] mismatch: got %f, want %f", i, resp.FFloats[i], v)
			}
		}
	}
}
