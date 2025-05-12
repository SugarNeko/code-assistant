package grpcbin_test

import (
	"context"
	"testing"
	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test",
		FStrings: []string{"test1", "test2"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   123456789,
		FInt64S:  []int64{12, 34, 56},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   1.23,
		FFloats:  []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp == nil {
		t.Fatalf("Expected non-nil response")
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString: %v, got: %v", req.FString, resp.FString)
	}
}
