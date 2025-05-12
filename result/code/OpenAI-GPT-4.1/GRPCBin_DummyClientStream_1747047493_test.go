package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyClientStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create client stream: %v", err)
	}

	var lastMessage *pb.DummyMessage
	for i := 1; i <= 10; i++ {
		msg := &pb.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"foo", "bar", "baz"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2, 3, int32(i)},
			FEnum:    pb.DummyMessage_ENUM_1,
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2},
			FSub:     &pb.DummyMessage_Sub{FString: "sub_string"},
			FSubs: []*pb.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:   true,
			FBools:  []bool{true, false},
			FInt64:  int64(i),
			FInt64S: []int64{int64(i), 42, 100},
			FBytes:  []byte("bytes"),
			FBytess: [][]byte{[]byte("A"), []byte("B")},
			FFloat:  float32(i) + 0.1,
			FFloats: []float32{1.1, 2.2, float32(i) + 0.1},
		}
		lastMessage = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send stream message %d: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive server response: %v", err)
	}

	// Validate the response - should match the last message sent
	if reply.FString != lastMessage.FString ||
		reply.FInt32 != lastMessage.FInt32 ||
		reply.FEnum != lastMessage.FEnum ||
		reply.FBool != lastMessage.FBool ||
		reply.FInt64 != lastMessage.FInt64 ||
		reply.FFloat != lastMessage.FFloat {
		t.Errorf("Server response does not match the last sent message. Got: %+v, Want: %+v", reply, lastMessage)
	}

	// Optionally, more detailed checks of slices, nested and repeated fields
	if len(reply.FStrings) != len(lastMessage.FStrings) {
		t.Errorf("Expected FStrings length %d, got %d", len(lastMessage.FStrings), len(reply.FStrings))
	}
}
