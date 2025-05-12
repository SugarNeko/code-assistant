package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
	"google.golang.org/protobuf/proto"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	req := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"a", "b"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "world"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "foo"}, {FString: "bar"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    12345,
		FInt64S:   []int64{111, 222},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 1.41},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive: %v", err)
	}

	// Validate: the server is expected to echo the request
	if !proto.Equal(resp, req) {
		t.Errorf("Response does not match request.\nGot:  %+v\nWant: %+v", resp, req)
	}

	// Close stream
	if err := stream.CloseSend(); err != nil {
		t.Errorf("Failed to close send: %v", err)
	}
}
