package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	sendMsg := &pb.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"foo", "bar"},
		FInt32:    12345,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    9876543210,
		FInt64S:   []int64{100, 200, 300},
		FBytes:    []byte("hello"),
		FBytess:   [][]byte{[]byte("foo"), []byte("bar")},
		FFloat:    3.1415,
		FFloats:   []float32{1.0, 2.0, 3.0},
	}

	if err := stream.Send(sendMsg); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	recvMsg, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Validate that sent and received messages are equal (basic fields)
	if sendMsg.FString != recvMsg.FString ||
		sendMsg.FInt32 != recvMsg.FInt32 ||
		sendMsg.FEnum != recvMsg.FEnum ||
		sendMsg.FSub.FString != recvMsg.FSub.FString ||
		sendMsg.FBool != recvMsg.FBool ||
		sendMsg.FInt64 != recvMsg.FInt64 ||
		sendMsg.FFloat != recvMsg.FFloat {
		t.Errorf("Mismatch between sent and received top-level fields")
	}

	if len(sendMsg.FStrings) != len(recvMsg.FStrings) {
		t.Errorf("Mismatch in FStrings length")
	}
	for i := range sendMsg.FStrings {
		if sendMsg.FStrings[i] != recvMsg.FStrings[i] {
			t.Errorf("Mismatch in FStrings[%d]: %s != %s", i, sendMsg.FStrings[i], recvMsg.FStrings[i])
		}
	}
}
