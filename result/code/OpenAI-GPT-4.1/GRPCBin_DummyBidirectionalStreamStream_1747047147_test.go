package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	// Set up connection options with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Construct a positive DummyMessage request
	req := &grpcbin.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"s1", "s2"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub-1"}, {FString: "sub-2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    9001,
		FInt64S:   []int64{1001, 1002},
		FBytes:    []byte("bytes-test"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 1.41},
	}

	// Send the request
	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive the response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Client response validation
	if resp.FString != req.FString {
		t.Errorf("FString: got %q, want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings: length got %d, want %d", len(resp.FStrings), len(req.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32: got %v, want %v", resp.FInt32, req.FInt32)
	}
	// ... Validate more fields as needed ...

	// (Optional) Close the stream from client side
	if err := stream.CloseSend(); err != nil {
		t.Errorf("Failed to close stream: %v", err)
	}
}
