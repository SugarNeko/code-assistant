package grpcbin_test

import (
	"context"
	"testing"
	"time"
	
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.DummyMessage{
		FString:  "hello",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
		FSub:     &pb.DummyMessage_Sub{FString: "subvalue"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:    true,
		FBools:   []bool{false, true},
		FInt64:   123456789,
		FInt64S:  []int64{9, 8, 7},
		FBytes:   []byte("abc"),
		FBytess:  [][]byte{[]byte("x"), []byte("y")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if !proto.Equal(resp, req) {
		t.Errorf("response does not match request.\nGot:  %#v\nWant: %#v", resp, req)
	}
}
