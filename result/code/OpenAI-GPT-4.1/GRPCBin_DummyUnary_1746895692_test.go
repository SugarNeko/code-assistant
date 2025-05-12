package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestGRPCBin_DummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"str1", "str2"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub_value"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    1000,
		FInt64S:   []int64{10, 20},
		FBytes:    []byte("foo"),
		FBytess:   [][]byte{[]byte("bar1"), []byte("bar2")},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 1.41},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary returned error: %v", err)
	}

	// Client-side validation: The response should echo the request (per service contract)
	if resp.FString != req.FString {
		t.Errorf("expected FString=%q, got %q", req.FString, resp.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("expected FStrings length=%d, got %d", len(req.FStrings), len(resp.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("expected FInt32=%d, got %d", req.FInt32, resp.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("expected FInt32S length=%d, got %d", len(req.FInt32S), len(resp.FInt32S))
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("expected FEnum=%v, got %v", req.FEnum, resp.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("expected FEnums length=%d, got %d", len(req.FEnums), len(resp.FEnums))
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("expected FSub.FString=%q, got %v", req.FSub.FString, resp.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("expected FSubs length=%d, got %d", len(req.FSubs), len(resp.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("expected FBool=%v, got %v", req.FBool, resp.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("expected FBools length=%d, got %d", len(req.FBools), len(resp.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("expected FInt64=%d, got %d", req.FInt64, resp.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("expected FInt64S length=%d, got %d", len(req.FInt64S), len(resp.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("expected FBytes=%q, got %q", req.FBytes, resp.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("expected FBytess length=%d, got %d", len(req.FBytess), len(resp.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("expected FFloat=%v, got %v", req.FFloat, resp.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("expected FFloats length=%d, got %d", len(req.FFloats), len(resp.FFloats))
	}
}
