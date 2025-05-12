package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req := &grpcbin.DummyMessage{
		FString:   "test",
		FStrings:  []string{"a", "b"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "a"}, {FString: "b"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    99,
		FInt64S:   []int64{7, 8, 9},
		FBytes:    []byte{0x1, 0x2},
		FBytess:   [][]byte{{0xA}, {0xB}},
		FFloat:    1.23,
		FFloats:   []float32{4.56, 7.89},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream RPC failed: %v", err)
	}

	recvCount := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		recvCount++

		// Validate main fields being echoed back
		if resp.FString != req.FString {
			t.Errorf("FString expected %q, got %q", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32 expected %d, got %d", req.FInt32, resp.FInt32)
		}
		if resp.FBool != req.FBool {
			t.Errorf("FBool expected %v, got %v", req.FBool, resp.FBool)
		}
	}

	if recvCount != 10 {
		t.Errorf("Expected 10 streamed responses, got %d", recvCount)
	}
}
