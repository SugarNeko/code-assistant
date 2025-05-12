package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	// Set up context with 15 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Connect to server
	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct a valid DummyMessage
	req := &grpcbin.DummyMessage{
		FString:   "hello world",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3, 4},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_1},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-message"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    9876543210,
		FInt64S:   []int64{11, 22, 33},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 1.62},
	}

	// Call the DummyServerStream rpc
	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	// Expecting a server stream of 10 identical DummyMessage copies
	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			if grpc.Code(err) == grpc.Code(context.Canceled) || grpc.Code(err) == grpc.Code(context.DeadlineExceeded) {
				t.Fatalf("Stream Recv interrupted: %v", err)
			}
			break // expected end of stream
		}
		count++

		// Validate the server response fields match our request
		if resp.FString != req.FString {
			t.Errorf("Response FString mismatch, got %q, want %q", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Response FInt32 mismatch, got %d, want %d", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Response FEnum mismatch, got %v, want %v", resp.FEnum, req.FEnum)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("Response FBytes mismatch, got %v, want %v", resp.FBytes, req.FBytes)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("Response FFloat mismatch, got %v, want %v", resp.FFloat, req.FFloat)
		}
	}

	if count != 10 {
		t.Errorf("Expected 10 streamed responses, got %d", count)
	}
}
