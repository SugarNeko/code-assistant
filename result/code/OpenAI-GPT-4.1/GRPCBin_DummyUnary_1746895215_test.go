package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct a valid DummyMessage request
	req := &grpcbin.DummyMessage{
		FString:  "hello",
		FStrings: []string{"one", "two"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub string"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "a"}, {FString: "b"}},
		FBool:    true,
		FBools:   []bool{true, false, true},
		FInt64:   9876543210,
		FInt64S:  []int64{7, 8, 9},
		FBytes:   []byte("test-bytes"),
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	// Client Response Validation: Check echoed fields
	if resp.FString != req.FString {
		t.Errorf("expected FString %q, got %q", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("expected FStrings len %d, got %d", len(req.FStrings), len(resp.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("expected FInt32 %d, got %d", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("expected FInt32S len %d, got %d", len(req.FInt32S), len(resp.FInt32S))
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("expected FEnum %d, got %d", req.FEnum, resp.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("expected FEnums len %d, got %d", len(req.FEnums), len(resp.FEnums))
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("expected FSub.FString %q, got %v", req.FSub.FString, resp.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("expected FSubs len %d, got %d", len(req.FSubs), len(resp.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("expected FBool %v, got %v", req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("expected FBools len %d, got %d", len(req.FBools), len(resp.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("expected FInt64 %d, got %d", req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("expected FInt64S len %d, got %d", len(req.FInt64S), len(resp.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("expected FBytes %q, got %q", req.FBytes, resp.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("expected FBytess len %d, got %d", len(req.FBytess), len(resp.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("expected FFloat %f, got %f", req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("expected FFloats len %d, got %d", len(req.FFloats), len(resp.FFloats))
	}

	// Additional validation as needed (e.g., check values in slices, etc.)
}
