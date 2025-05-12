package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test_string",
		FStrings: []string{"a", "b", "c"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_2,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "sub_string"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "list1"}, {FString: "list2"}},
		FBool:    true,
		FBools:   []bool{false, true, true},
		FInt64:   9999999999,
		FInt64S:  []int64{7, 8, 9},
		FBytes:   []byte("bytes-value"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if !proto.Equal(resp, req) {
		t.Errorf("DummyUnary expected response to echo request, want=%v, got=%v", req, resp)
	}
}
