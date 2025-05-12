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
		FStrings: []string{"foo", "bar"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "subfield"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   321,
		FInt64S:  []int64{101, 202},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{[]byte("a"), []byte("b")},
		FFloat:   1.23,
		FFloats:  []float32{4.56, 7.89},
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
		// Validate server response matches expected format (echoed message)
		if resp.FString != req.FString {
			t.Errorf("Response FString = %q; want %q", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Response FInt32 = %d; want %d", resp.FInt32, req.FInt32)
		}
		// ... add more validation as needed for each field
		count++
	}

	if count != 10 {
		t.Errorf("expected 10 streamed responses, got %d", count)
	}
}
