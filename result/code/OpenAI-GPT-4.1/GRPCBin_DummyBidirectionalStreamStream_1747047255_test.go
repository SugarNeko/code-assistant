package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	req := &pb.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"one", "two"},
		FInt32:    42,
		FInt32S:   []int32{7, 8},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub1"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub2"}, {FString: "sub3"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    99,
		FInt64S:   []int64{123, 456},
		FBytes:    []byte{0x10, 0x20},
		FBytess:   [][]byte{{0x1, 0x2}, {0x3, 0x4}},
		FFloat:    1.23,
		FFloats:   []float32{4.56, 7.89},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("failed to send request message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	// Validate response: Should echo back the request
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %v, want %v", len(resp.FStrings), len(req.FStrings))
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
	if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %q, want %q", resp.FSub.FString, req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %v, want %v", len(resp.FSubs), len(req.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %v, want %v", len(resp.FBools), len(req.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %v, want %v", len(resp.FInt64S), len(req.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %v, want %v", len(resp.FBytess), len(req.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %v, want %v", len(resp.FFloats), len(req.FFloats))
	}

	// Optionally, close stream
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("failed to close stream: %v", err)
	}
}
