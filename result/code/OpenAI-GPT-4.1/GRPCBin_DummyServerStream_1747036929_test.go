package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream_Positive(t *testing.T) {
	addr := "grpcb.in:9000"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"one", "two"},
		FInt32:   123,
		FInt32S:  []int32{11, 12},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-str"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   9876543210,
		FInt64S:  []int64{123456789, 987654321},
		FBytes:   []byte("hello"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   1.234,
		FFloats:  []float32{2.34, 5.67},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("stream.Recv failed at msg %d: %v", i, err)
		}

		// Validate each response field equals sent request
		if resp.FString != req.FString {
			t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
		}
		if len(resp.FInt32S) != len(req.FInt32S) {
			t.Errorf("FInt32S length mismatch: got %d, want %d", len(resp.FInt32S), len(req.FInt32S))
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if len(resp.FEnums) != len(req.FEnums) {
			t.Errorf("FEnums length mismatch: got %d, want %d", len(resp.FEnums), len(req.FEnums))
		}
		if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
			t.Errorf("FSub mismatch: got %+v, want %+v", resp.FSub, req.FSub)
		}
		if len(resp.FSubs) != len(req.FSubs) {
			t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
		}
		if resp.FBool != req.FBool {
			t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
		}
		if len(resp.FBools) != len(req.FBools) {
			t.Errorf("FBools length mismatch: got %d, want %d", len(resp.FBools), len(req.FBools))
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, req.FInt64)
		}
		if len(resp.FInt64S) != len(req.FInt64S) {
			t.Errorf("FInt64S length mismatch: got %d, want %d", len(resp.FInt64S), len(req.FInt64S))
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("FBytes mismatch: got %q, want %q", resp.FBytes, req.FBytes)
		}
		if len(resp.FBytess) != len(req.FBytess) {
			t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, req.FFloat)
		}
		if len(resp.FFloats) != len(req.FFloats) {
			t.Errorf("FFloats length mismatch: got %d, want %d", len(resp.FFloats), len(req.FFloats))
		}
	}
	// End of stream should be EOF
	_, err = stream.Recv()
	if err == nil {
		t.Errorf("Expected EOF from stream.Recv after all messages, got nil error")
	}
}
