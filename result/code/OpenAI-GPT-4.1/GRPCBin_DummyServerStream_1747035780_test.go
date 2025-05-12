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
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"str1", "str2"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false, true},
		FInt64:   987654321,
		FInt64S:  []int64{100, 200, 300},
		FBytes:   []byte("bytes_data"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed to start: %v", err)
	}

	msgCount := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		msgCount++

		// Validate fields match the request
		if resp.FString != req.FString {
			t.Errorf("response FString mismatch: got=%v want=%v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response FInt32 mismatch: got=%v want=%v", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("response FEnum mismatch: got=%v want=%v", resp.FEnum, req.FEnum)
		}
		if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
			t.Errorf("response FSub.FString mismatch: got=%v want=%v", resp.FSub, req.FSub)
		}
		if !resp.FBool == req.FBool {
			t.Errorf("response FBool mismatch: got=%v want=%v", resp.FBool, req.FBool)
		}
		// Can add further deep validation per field/repeated fields as needed
	}

	if msgCount != 10 {
		t.Errorf("Expected 10 messages in stream, got %d", msgCount)
	}
}
