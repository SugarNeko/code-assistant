package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyServerStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"a", "b", "c"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    1234567,
		FInt64S:   []int64{123, 456},
		FBytes:    []byte{0x01, 0x02},
		FBytess:   [][]byte{{0x0a, 0x0b}, {0x0c}},
		FFloat:    1.2345,
		FFloats:   []float32{9.8, 7.6},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	var recvCount int
	for {
		resp, recvErr := stream.Recv()
		if recvErr != nil {
			if recvErr.Error() == "EOF" {
				break
			}
			t.Fatalf("DummyServerStream.Recv() failed: %v", recvErr)
		}
		recvCount++

		// Validate server response content matches sent request (should echo it 10 times per spec)
		if resp.FString != req.FString {
			t.Errorf("FString mismatch. got %q, want %q", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 mismatch. got %d, want %d", resp.FInt32, req.FInt32)
		}
		// (Do similar checks for other important fields if needed)
	}

	if recvCount != 10 {
		t.Errorf("expected 10 messages, got %d", recvCount)
	}
}
