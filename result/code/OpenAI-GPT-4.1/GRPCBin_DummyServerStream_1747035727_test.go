package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"a", "b", "c"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub-str"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    1002,
		FInt64S:   []int64{123, 456},
		FBytes:    []byte("hello"),
		FBytess:   [][]byte{[]byte("foo"), []byte("bar")},
		FFloat:    1.23,
		FFloats:   []float32{2.34, 3.45},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream initiation failed: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		count++

		// Client response validation
		if resp.FString != req.FString {
			t.Errorf("FString mismatch: got %q, want %q", resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if resp.FBool != req.FBool {
			t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
		}
	}
	if count != 10 {
		t.Errorf("expected 10 messages, got %d", count)
	}
}
