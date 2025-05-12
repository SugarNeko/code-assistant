package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test string",
		FStrings:  []string{"one", "two"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
		FSub:      &pb.DummyMessage_Sub{FString: "subtest"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "a"}, {FString: "b"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    987654321,
		FInt64S:   []int64{11, 22},
		FBytes:    []byte{0xCA, 0xFE, 0xBA, 0xBE},
		FBytess:   [][]byte{{0xDE, 0xAD}, {0xBE, 0xEF}},
		FFloat:    3.1415,
		FFloats:   []float32{2.71, 1.41},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	recvCount := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		recvCount++

		// Response validation: ensure it mirrors the request
		if resp.FString != req.FString {
			t.Errorf("Response FString mismatch: got %v, want %v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Response FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Response FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if !resp.FBool == req.FBool {
			t.Errorf("Response FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
		}
	}

	if recvCount != 10 {
		t.Errorf("Expected 10 streamed responses, got %d", recvCount)
	}
}
