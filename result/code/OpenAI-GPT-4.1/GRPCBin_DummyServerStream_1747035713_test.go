package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	// Set up gRPC connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct a valid DummyMessage following the spec
	req := &grpcbin.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"a", "b", "c"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-str"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    123456789,
		FInt64S:   []int64{10, 20, 30},
		FBytes:    []byte("hello-bytes"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	// We expect 10 responses with the same content as req
	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Error receiving on stream at %d: %v", i, err)
		}

		// Validate individual fields for response
		if resp.FString != req.FString {
			t.Errorf("FString mismatch: got %v want %v", resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("FStrings length mismatch: got %d want %d", len(resp.FStrings), len(req.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 mismatch: got %v want %v", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("FEnum mismatch: got %v want %v", resp.FEnum, req.FEnum)
		}
		if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
			t.Errorf("FSub FString mismatch: got %v want %v", resp.FSub, req.FSub)
		}
		if resp.FBool != req.FBool {
			t.Errorf("FBool mismatch: got %v want %v", resp.FBool, req.FBool)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("FInt64 mismatch: got %v want %v", resp.FInt64, req.FInt64)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("FBytes mismatch: got %v want %v", resp.FBytes, req.FBytes)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("FFloat mismatch: got %v want %v", resp.FFloat, req.FFloat)
		}
		// (Add additional array/slice/multi-element checks if thorough validation is needed.)
	}
	// Ensure the stream closes correctly after 10 messages
	_, err = stream.Recv()
	if err == nil {
		t.Errorf("Expected end of stream after 10 messages, but got more data")
	}
}
