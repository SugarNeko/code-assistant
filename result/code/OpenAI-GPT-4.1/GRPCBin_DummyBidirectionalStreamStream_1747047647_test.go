package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("DummyBidirectionalStreamStream failed to start: %v", err)
	}

	req := &pb.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"one", "two"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_2,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "subfield"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    9876543210,
		FInt64S:   []int64{111, 222},
		FBytes:    []byte("bytedata"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.1415,
		FFloats:   []float32{1.1, 2.2},
	}

	// Send
	if err := stream.Send(req); err != nil {
		t.Fatalf("failed to send DummyMessage: %v", err)
	}

	// Receive
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive DummyMessage: %v", err)
	}

	// Validate response matches request fields
	if resp.FString != req.FString {
		t.Errorf("response FString mismatch: got %q want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("response FStrings length mismatch: got %v want %v", len(resp.FStrings), len(req.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("response FInt32 mismatch: got %d want %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("response FInt32S length mismatch: got %v want %v", len(resp.FInt32S), len(req.FInt32S))
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("response FEnum mismatch: got %v want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("response FEnums length mismatch: got %v want %v", len(resp.FEnums), len(req.FEnums))
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("response FSub.FString mismatch: got %v want %v", resp.FSub, req.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("response FSubs length mismatch: got %v want %v", len(resp.FSubs), len(req.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("response FBool mismatch: got %v want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("response FBools length mismatch: got %v want %v", len(resp.FBools), len(req.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("response FInt64 mismatch: got %v want %v", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("response FInt64S length mismatch: got %v want %v", len(resp.FInt64S), len(req.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("response FBytes mismatch: got %v want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("response FBytess length mismatch: got %v want %v", len(resp.FBytess), len(req.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("response FFloat mismatch: got %v want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("response FFloats length mismatch: got %v want %v", len(resp.FFloats), len(req.FFloats))
	}
}
