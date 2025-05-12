package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	// Dial with 15s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("DummyBidirectionalStreamStream creation failed: %v", err)
	}

	// Construct a valid DummyMessage
	msg := &pb.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"foo", "bar"},
		FInt32:    123,
		FInt32S:   []int32{456, 789},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
		FSub:      &pb.DummyMessage_Sub{FString: "nested"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    9876543210,
		FInt64S:   []int64{111, 222},
		FBytes:    []byte("bytes-demo"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    1.23,
		FFloats:   []float32{3.14, 2.71},
	}

	// Send a message
	if err := stream.Send(msg); err != nil {
		t.Fatalf("failed to send: %v", err)
	}

	// Receive echo response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive: %v", err)
	}

	// Validate response equals sent message
	if resp.FString != msg.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, msg.FString)
	}
	if resp.FInt32 != msg.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, msg.FInt32)
	}
	if resp.FEnum != msg.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, msg.FEnum)
	}
	if resp.FSub == nil || resp.FSub.FString != msg.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %q, want %q", resp.FSub.FString, msg.FSub.FString)
	}
	if resp.FBool != msg.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, msg.FBool)
	}
	if resp.FInt64 != msg.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, msg.FInt64)
	}
	if string(resp.FBytes) != string(msg.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, msg.FBytes)
	}
	if resp.FFloat != msg.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, msg.FFloat)
	}

	// Optional: Validate slices for representative fields
	for i := range msg.FStrings {
		if i >= len(resp.FStrings) || resp.FStrings[i] != msg.FStrings[i] {
			t.Errorf("FStrings[%d] mismatch: got %q, want %q", i, resp.FStrings[i], msg.FStrings[i])
		}
	}
	for i := range msg.FInts {
		if i >= len(resp.FInts) || resp.FInts[i] != msg.FInts[i] {
			t.Errorf("FInts[%d] mismatch: got %v, want %v", i, resp.FInts[i], msg.FInts[i])
		}
	}

	// Close client side stream
	if err := stream.CloseSend(); err != nil {
		t.Errorf("failed to close stream send: %v", err)
	}
}
