package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	// Set up a connection to the server with 15 seconds timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"a", "b"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "foo"}, {FString: "bar"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    123456789,
		FInt64S:   []int64{11, 22},
		FBytes:    []byte{0x1, 0x2, 0x3},
		FBytess:   [][]byte{{0x11, 0x22}, {0x33}},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 1.61},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	expectedCount := 10
	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		// Validate response matches the request data (echoed back)
		if resp.FString != req.FString {
			t.Errorf("FString: got %q, want %q", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("FInt32: got %v, want %v", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("FEnum: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if resp.FBool != req.FBool {
			t.Errorf("FBool: got %v, want %v", resp.FBool, req.FBool)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("FInt64: got %v, want %v", resp.FInt64, req.FInt64)
		}
		// You may add further checks for slices and nested messages as needed.

		count++
	}

	if count != expectedCount {
		t.Errorf("Expected %d responses, got %d", expectedCount, count)
	}
}
