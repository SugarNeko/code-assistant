package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "hello",
		FStrings: []string{"one", "two"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   9999,
		FInt64S:  []int64{100, 200},
		FBytes:   []byte("abc"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   1.23,
		FFloats:  []float32{2.34, 3.45},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		count++

		// Validate response
		if resp.FString != req.FString {
			t.Errorf("expected FString %q, got %q", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("expected FInt32 %d, got %d", req.FInt32, resp.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("expected FEnum %v, got %v", req.FEnum, resp.FEnum)
		}
	}

	if count != 10 {
		t.Errorf("expected 10 messages, got %d", count)
	}
}
