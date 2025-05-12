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

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "example",
		FStrings:  []string{"str1", "str2"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    987654321,
		FInt64S:   []int64{111, 222, 333},
		FBytes:    []byte("bytesVal"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	// The server should send back 10 times the received DummyMessage
	resultCount := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		resultCount++

		// Validate select response fields match what we sent
		if resp.FString != req.FString {
			t.Errorf("response FString got %v want %v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response FInt32 got %v want %v", resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("response FEnum got %v want %v", resp.FEnum, req.FEnum)
		}
		if resp.FBool != req.FBool {
			t.Errorf("response FBool got %v want %v", resp.FBool, req.FBool)
		}
	}

	if resultCount != 10 {
		t.Errorf("expected 10 streamed DummyMessages, got %d", resultCount)
	}
}
