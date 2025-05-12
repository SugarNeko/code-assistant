package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin" // Adjust import path if necessary

	"google.golang.org/protobuf/proto"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to open stream: %v", err)
	}

	req := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_2,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_0},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   987654321,
		FInt64S:  []int64{100, 200, 300},
		FBytes:   []byte("hello"),
		FBytess:  [][]byte{[]byte("a"), []byte("b")},
		FFloat:   3.14,
		FFloats:  []float32{2.71, 1.61},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	recv, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive: %v", err)
	}

	if !proto.Equal(req, recv) {
		t.Errorf("Received response does not match sent request\nSent:    %+v\nReceived:%+v", req, recv)
	}
}
