package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("failed to connect to grpcb.in:9000: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"a", "b"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "substring"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "x"}, {FString: "y"}},
		FBool:    true,
		FBools:   []bool{false, true},
		FInt64:   999,
		FInt64S:  []int64{100, 200},
		FBytes:   []byte("abc"),
		FBytess:  [][]byte{[]byte("foo"), []byte("bar")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream RPC failed: %v", err)
	}

	respCount := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		respCount++

		// Validate that the response matches the sent request.
		if resp.FString != req.FString {
			t.Errorf("response FString = %q; want %q", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response FInt32 = %d; want %d", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("response FEnum = %v; want %v", resp.FEnum, req.FEnum)
		}
		// (Further field-by-field comparison can be added here)
	}

	if respCount != 10 {
		t.Errorf("expected 10 responses from DummyServerStream, got %d", respCount)
	}
}
