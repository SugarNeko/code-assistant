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

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("DummyBidirectionalStreamStream failed: %v", err)
	}

	// Constructing a typical DummyMessage request
	testMsg := &pb.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"one", "two"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_2,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_1},
		FSub:      &pb.DummyMessage_Sub{FString: "sub_test"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub_1"}, {FString: "sub_2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    1234567890,
		FInt64S:   []int64{100, 200},
		FBytes:    []byte("byte-data"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    12.34,
		FFloats:   []float32{1.23, 4.56},
	}

	// Send the request
	if err := stream.Send(testMsg); err != nil {
		t.Fatalf("Send() failed: %v", err)
	}

	// Receive the echoed response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Recv() failed: %v", err)
	}

	// Client response validation
	if resp.FString != testMsg.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, testMsg.FString)
	}
	if len(resp.FStrings) != len(testMsg.FStrings) {
		t.Errorf("FStrings len mismatch: got %d, want %d", len(resp.FStrings), len(testMsg.FStrings))
	}
	if resp.FInt32 != testMsg.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, testMsg.FInt32)
	}
	if len(resp.FInt32S) != len(testMsg.FInt32S) {
		t.Errorf("FInt32S len mismatch: got %d, want %d", len(resp.FInt32S), len(testMsg.FInt32S))
	}
	if resp.FEnum != testMsg.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, testMsg.FEnum)
	}
	if len(resp.FEnums) != len(testMsg.FEnums) {
		t.Errorf("FEnums len mismatch: got %d, want %d", len(resp.FEnums), len(testMsg.FEnums))
	}
	if resp.FSub == nil || testMsg.FSub == nil || resp.FSub.FString != testMsg.FSub.FString {
		t.Errorf("FSub FString mismatch: got %q, want %q", resp.FSub.FString, testMsg.FSub.FString)
	}
	if len(resp.FSubs) != len(testMsg.FSubs) {
		t.Errorf("FSubs len mismatch: got %d, want %d", len(resp.FSubs), len(testMsg.FSubs))
	}
	if resp.FBool != testMsg.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, testMsg.FBool)
	}
	if len(resp.FBools) != len(testMsg.FBools) {
		t.Errorf("FBools len mismatch: got %d, want %d", len(resp.FBools), len(testMsg.FBools))
	}
	if resp.FInt64 != testMsg.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, testMsg.FInt64)
	}
	if len(resp.FInt64S) != len(testMsg.FInt64S) {
		t.Errorf("FInt64S len mismatch: got %d, want %d", len(resp.FInt64S), len(testMsg.FInt64S))
	}
	if string(resp.FBytes) != string(testMsg.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, testMsg.FBytes)
	}
	if len(resp.FBytess) != len(testMsg.FBytess) {
		t.Errorf("FBytess len mismatch: got %d, want %d", len(resp.FBytess), len(testMsg.FBytess))
	}
	if resp.FFloat != testMsg.FFloat {
		t.Errorf("FFloat mismatch: got %f, want %f", resp.FFloat, testMsg.FFloat)
	}
	if len(resp.FFloats) != len(testMsg.FFloats) {
		t.Errorf("FFloats len mismatch: got %d, want %d", len(resp.FFloats), len(testMsg.FFloats))
	}
}
