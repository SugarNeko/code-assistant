package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	address := "grpcb.in:9000"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req := &grpcbin.DummyMessage{
		FString:  "hello",
		FStrings: []string{"foo", "bar"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub1"},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{FString: "sub2"},
			{FString: "sub3"},
		},
		FBool:    true,
		FBools:   []bool{true, false, true},
		FInt64:   9876543210,
		FInt64S:  []int64{111, 222},
		FBytes:   []byte("bytes data"),
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	expectedMsgCount := 10
	for i := 0; i < expectedMsgCount; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("stream.Recv() failed at message %d: %v", i, err)
		}
		// Client response validation: check some fields
		if resp.FString != req.FString {
			t.Errorf("response FString mismatch at message %d: got %q, want %q", i, resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response FInt32 mismatch at message %d: got %d, want %d", i, resp.FInt32, req.FInt32)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("response FStrings length mismatch at message %d: got %d, want %d", i, len(resp.FStrings), len(req.FStrings))
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("response FEnum mismatch at message %d: got %v, want %v", i, resp.FEnum, req.FEnum)
		}
		if resp.FBool != req.FBool {
			t.Errorf("response FBool mismatch at message %d: got %v, want %v", i, resp.FBool, req.FBool)
		}
	}

	_, err = stream.Recv()
	if err == nil {
		t.Errorf("expected stream to be closed after %d messages, but got more messages", expectedMsgCount)
	}
}
