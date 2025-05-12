package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	proto "code-assistant/proto/grpcbin"
)

func TestGRPCBin(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGRPCBinClient(conn)

	// Positive testing
	req := &proto.DummyMessage{
		FString: "Hello",
		FStrings: []string{
			"String1",
			"String2",
		},
		FInt32: 123,
		FInt32s: []int32{
			10,
			20,
		},
		FPSub: &proto.DummyMessage_Sub{
			FString: "Sub",
		},
		FSubs: []*proto.DummyMessage_Sub{
			{FString: "Sub1"},
			{FString: "Sub2"},
		},
		FBool: true,
		FBools: []bool{
			true,
			false,
		},
		FInt64: 1234567890,
		FInt64s: []int64{
			1234567890,
			9876543210,
		},
		FBytes: []byte{1, 2, 3, 4},
		FBytess: [][]byte{
			{1, 2},
			{3, 4},
		},
		FFloat: 3.14,
		FFloats: []float32{
			3.14,
			2.71,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Errorf("client.DummyUnary(_) = _, %v; want nil", err)
		return
	}

	if resp == nil {
		t.Errorf("client.DummyUnary(_) = nil, nil; want non-nil response")
		return
	}

	if resp.FString != req.FString {
		t.Errorf("resp.FString = %q, want %q", resp.FString, req.FString)
	}

	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("resp.FStrings = %v, want %v", resp.FStrings, req.FStrings)
	}

	for i, v := range resp.FStrings {
		if v != req.FStrings[i] {
			t.Errorf("resp.FStrings[%d] = %q, want %q", i, v, req.FStrings[i])
		}
	}
}
