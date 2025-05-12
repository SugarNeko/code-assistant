package grpcbin_test

import (
	"context"
	"testing"
	"time"
	"io"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
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
		FSub:      &grpcbin.DummyMessage_Sub{FString: "substr"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    123456789,
		FInt64S:   []int64{11, 22},
		FBytes:    []byte("bytesample"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    1.23,
		FFloats:   []float32{2.34, 5.67},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream call failed: %v", err)
	}

	received := 0
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Error receiving from server: %v", err)
		}

		// --- Client Response Validation ---
		if resp.FString != req.FString {
			t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
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
			t.Errorf("FSub.FString mismatch: got %v, want %v", resp.FSub, req.FSub)
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
			t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
		}
		if len(resp.FInt64S) != len(req.FInt64S) {
			t.Errorf("FInt64S length mismatch: got %d, want %d", len(resp.FInt64S), len(req.FInt64S))
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
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

		received++
	}

	// --- Server Response Validation ---
	if received != 10 {
		t.Errorf("Expected 10 streamed responses, got %d", received)
	}
}
