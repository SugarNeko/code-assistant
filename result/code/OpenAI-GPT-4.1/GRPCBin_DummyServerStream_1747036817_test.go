package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "subvalue"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{false, true, false},
		FInt64:    1001,
		FInt64S:   []int64{2002, 3003},
		FBytes:    []byte("bytestring"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream error: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		count++

		// Client response validation: check response fields equal to req fields.
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
			t.Errorf("expected FEnum %v, got %v", req.FEnum, resp.FEnum)
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
			t.Errorf("expected FBytes %q, got %q", string(req.FBytes), string(resp.FBytes))
		}
		if len(resp.FBytess) != len(req.FBytess) {
			t.Errorf("expected FBytess len %d, got %d", len(req.FBytess), len(resp.FBytess))
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("expected FFloat %v, got %v", req.FFloat, resp.FFloat)
		}
		if len(resp.FFloats) != len(req.FFloats) {
			t.Errorf("expected FFloats len %d, got %d", len(req.FFloats), len(resp.FFloats))
		}
	}

	// Test server response validation: make sure exactly 10 responses were returned.
	if count != 10 {
		t.Errorf("expected 10 responses from server, got %d", count)
	}
}
