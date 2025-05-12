package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "sub-1"}, {FString: "sub-2"}},
		FBool:    true,
		FBools:   []bool{true, false, true},
		FInt64:   64,
		FInt64S:  []int64{100, 200},
		FBytes:   []byte("hello-bytes"),
		FBytess:  [][]byte{[]byte("a"), []byte("b")},
		FFloat:   3.14,
		FFloats:  []float32{6.28, 1.414},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream RPC failed: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		count++
		if resp.FString != req.FString {
			t.Errorf("Response FString mismatch: got %v, want %v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Response FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Response FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
		}
		// Optionally add more field validations as needed
	}
	if count != 10 {
		t.Errorf("Expected 10 streamed messages, got %d", count)
	}
}
