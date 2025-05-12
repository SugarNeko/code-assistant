package grpcbin_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:    "test",
		FStrings:   []string{"a", "b"},
		FInt32:     42,
		FInt32S:    []int32{1, 2},
		FEnum:      grpcbin.DummyMessage_ENUM_1,
		FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:       &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:      true,
		FBools:     []bool{true, false},
		FInt64:     123456789,
		FInt64S:    []int64{987654321, 123456789},
		FBytes:     []byte("test bytes"),
		FBytess:    [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:     3.14,
		FFloats:    []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
	}
	if !reflect.DeepEqual(resp.FStrings, req.FStrings) {
		t.Errorf("FStrings mismatch: got %v, want %v", resp.FStrings, req.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
	}
	if !reflect.DeepEqual(resp.FInt32S, req.FInt32S) {
		t.Errorf("FInt32S mismatch: got %v, want %v", resp.FInt32S, req.FInt32S)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if !reflect.DeepEqual(resp.FEnums, req.FEnums) {
		t.Errorf("FEnums mismatch: got %v, want %v", resp.FEnums, req.FEnums)
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %q, want %q", resp.FSub.FString, req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
	} else {
		for i := range resp.FSubs {
			if resp.FSubs[i].FString != req.FSubs[i].FString {
				t.Errorf("FSubs[%d].FString mismatch: got %q, want %q", i, resp.FSubs[i].FString, req.FSubs[i].FString)
			}
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %t, want %t", resp.FBool, req.FBool)
	}
	if !reflect.DeepEqual(resp.FBools, req.FBools) {
		t.Errorf("FBools mismatch: got %v, want %v", resp.FBools, req.FBools)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
	}
	if !reflect.DeepEqual(resp.FInt64S, req.FInt64S) {
		t.Errorf("FInt64S mismatch: got %v, want %v", resp.FInt64S, req.FInt64S)
	}
	if !reflect.DeepEqual(resp.FBytes, req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if !reflect.DeepEqual(resp.FBytess, req.FBytess) {
		t.Errorf("FBytess mismatch: got %v, want %v", resp.FBytess, req.FBytess)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %f, want %f", resp.FFloat, req.FFloat)
	}
	if !reflect.DeepEqual(resp.FFloats, req.FFloats) {
		t.Errorf("FFloats mismatch: got %v, want %v", resp.FFloats, req.FFloats)
	}
}
