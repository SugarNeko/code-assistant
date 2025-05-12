package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test string",
		FStrings: []string{"str1", "str2"},
		FInt32:   12345,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnumS:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "sub string"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   67890,
		FInt64S:  []int64{10, 20, 30},
		FBytes:   []byte("bytes"),
		FBytesS:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString: %v, got: %v", req.FString, resp.FString)
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32: %v, got: %v", req.FInt32, resp.FInt32)
	}

	// Add additional response validation as needed
}
