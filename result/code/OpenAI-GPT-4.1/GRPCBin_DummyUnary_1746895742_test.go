package grpcbin_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-hello"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    9876543210,
		FInt64S:   []int64{11, 22, 33},
		FBytes:    []byte("abcde"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    1.23,
		FFloats:   []float32{3.14, 33.3},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary returned error: %v", err)
	}

	// Server response validation (should echo the request)
	if resp.FString != req.FString {
		t.Errorf("Expected FString %q, got %q", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("Expected FStrings len %v, got %v", len(req.FStrings), len(resp.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32 %d, got %d", req.FInt32, resp.FInt32)
	}
	for i, v := range req.FInt32S {
		if i >= len(resp.FInt32S) || resp.FInt32S[i] != v {
			t.Errorf("Mismatch in FInt32S at index %d: expected %d, got %d", i, v, resp.FInt32S[i])
		}
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("Expected FEnum %v, got %v", req.FEnum, resp.FEnum)
	}
	for i, v := range req.FEnums {
		if i >= len(resp.FEnums) || resp.FEnums[i] != v {
			t.Errorf("Mismatch in FEnums at index %d: expected %v, got %v", i, v, resp.FEnums[i])
		}
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub mismatch: expected %q, got %+v", req.FSub.FString, resp.FSub)
	}
	for i, v := range req.FSubs {
		if i >= len(resp.FSubs) || resp.FSubs[i].FString != v.FString {
			t.Errorf("Mismatch in FSubs at index %d: expected %q, got %q", i, v.FString, resp.FSubs[i].FString)
		}
	}
	if resp.FBool != req.FBool {
		t.Errorf("Expected FBool %v, got %v", req.FBool, resp.FBool)
	}
	for i, v := range req.FBools {
		if i >= len(resp.FBools) || resp.FBools[i] != v {
			t.Errorf("Mismatch in FBools at index %d: expected %v, got %v", i, v, resp.FBools[i])
		}
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("Expected FInt64 %d, got %d", req.FInt64, resp.FInt64)
	}
	for i, v := range req.FInt64S {
		if i >= len(resp.FInt64S) || resp.FInt64S[i] != v {
			t.Errorf("Mismatch in FInt64S at index %d: expected %d, got %d", i, v, resp.FInt64S[i])
		}
	}
	if !bytes.Equal(resp.FBytes, req.FBytes) {
		t.Errorf("Expected FBytes %v, got %v", req.FBytes, resp.FBytes)
	}
	for i, v := range req.FBytess {
		if i >= len(resp.FBytess) || !bytes.Equal(resp.FBytess[i], v) {
			t.Errorf("Mismatch in FBytess at index %d: expected %v, got %v", i, v, resp.FBytess[i])
		}
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("Expected FFloat %f, got %f", req.FFloat, resp.FFloat)
	}
	for i, v := range req.FFloats {
		if i >= len(resp.FFloats) || resp.FFloats[i] != v {
			t.Errorf("Mismatch in FFloats at index %d: expected %f, got %f", i, v, resp.FFloats[i])
		}
	}
}
