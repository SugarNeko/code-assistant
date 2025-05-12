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
	// Setup connection with 15 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("DummyBidirectionalStreamStream failed to open stream: %v", err)
	}

	// Compose a positive (valid) DummyMessage
	msg := &pb.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"one", "two"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_2,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    12345678,
		FInt64S:   []int64{987654321, 19},
		FBytes:    []byte("bytes-data"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    1.618,
		FFloats:   []float32{2.718, 3.14},
	}

	// Send message
	if err := stream.Send(msg); err != nil {
		t.Fatalf("Failed to send stream request: %v", err)
	}

	// Receive response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive stream response: %v", err)
	}

	// Client response validation - echo semantics
	if resp.FString != msg.FString {
		t.Errorf("FString mismatch: got %s, want %s", resp.FString, msg.FString)
	}
	if resp.FInt32 != msg.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, msg.FInt32)
	}
	if resp.FEnum != msg.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, msg.FEnum)
	}
	if resp.FBool != msg.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, msg.FBool)
	}
	if resp.FInt64 != msg.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, msg.FInt64)
	}
	if string(resp.FBytes) != string(msg.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, msg.FBytes)
	}
	if resp.FFloat != msg.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, msg.FFloat)
	}
	if len(resp.FStrings) != len(msg.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(msg.FStrings))
	}
	if len(resp.FInt32S) != len(msg.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %d, want %d", len(resp.FInt32S), len(msg.FInt32S))
	}
	if len(resp.FEnums) != len(msg.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(resp.FEnums), len(msg.FEnums))
	}
	if resp.FSub == nil || msg.FSub == nil || resp.FSub.FString != msg.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v, want %v", resp.FSub, msg.FSub)
	}
	if len(resp.FSubs) != len(msg.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(msg.FSubs))
	}
	if len(resp.FBools) != len(msg.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(resp.FBools), len(msg.FBools))
	}
	if len(resp.FInt64S) != len(msg.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %d, want %d", len(resp.FInt64S), len(msg.FInt64S))
	}
	if len(resp.FBytess) != len(msg.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(msg.FBytess))
	}
	if len(resp.FFloats) != len(msg.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(resp.FFloats), len(msg.FFloats))
	}

	// Optionally, close the stream gracefully
	if err := stream.CloseSend(); err != nil {
		t.Errorf("Failed to close stream: %v", err)
	}
}
