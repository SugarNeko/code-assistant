package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
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
		FString:  "foo",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_2,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_0},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "substr"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:    true,
		FBools:   []bool{true, false, true},
		FInt64:   10000000000,
		FInt64S:  []int64{100, 200},
		FBytes:   []byte{0xAA, 0xBB},
		FBytess:  [][]byte{{0x01}, {0x02}},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	receivedCount := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		receivedCount++

		if resp.FString != req.FString {
			t.Errorf("Response FString got %v, want %v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Response FInt32 got %v, want %v", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Response FEnum got %v, want %v", resp.FEnum, req.FEnum)
		}
		// Further deep equals and checks for all fields can be added as necessary
	}

	if receivedCount != 10 {
		t.Errorf("Expected 10 responses from stream, got %d", receivedCount)
	}
}
