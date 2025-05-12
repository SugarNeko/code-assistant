package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "subfield"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "one"}, {FString: "two"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    1234567890,
		FInt64S:   []int64{100, 200, 300},
		FBytes:    []byte("byte-test"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("stream.Recv() failed: %v", err)
		}

		if resp.FString != req.FString {
			t.Errorf("response[%d].FString = %q, want %q", i, resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response[%d].FInt32 = %d, want %d", i, resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("response[%d].FEnum = %v, want %v", i, resp.FEnum, req.FEnum)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("response[%d].FBytes = %q, want %q", i, resp.FBytes, req.FBytes)
		}
	}

	// expect EOF after 10 messages
	_, err = stream.Recv()
	if err == nil {
		t.Errorf("expected EOF after 10 messages, got nil error")
	}
}
