package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "submsg"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    1234567890,
		FInt64S:   []int64{11, 22, 33},
		FBytes:    []byte("bytestr"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 42.0},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream error: %v", err)
	}

	var count int
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			t.Fatalf("Receiving from stream failed: %v", err)
		}
		count++

		// Server response validation - check for echo-like behavior
		if resp.FString != req.FString {
			t.Errorf("FString mismatch: got %v, want %v", resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
		}
		if len(resp.FInt32S) != len(req.FInt32S) {
			t.Errorf("FInt32S length mismatch: got %v, want %v", len(resp.FInt32S), len(req.FInt32S))
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if len(resp.FEnums) != len(req.FEnums) {
			t.Errorf("FEnums length mismatch: got %v, want %v", len(resp.FEnums), len(req.FEnums))
		}
		if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
			t.Errorf("FSub FString mismatch: got %+v, want %+v", resp.FSub, req.FSub)
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
	}
	if count != 10 {
		t.Errorf("Expected 10 streamed messages, got %d", count)
	}
}
