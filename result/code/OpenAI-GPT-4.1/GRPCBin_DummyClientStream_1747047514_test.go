package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
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

	client := grpcbin.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	messages := make([]*grpcbin.DummyMessage, 10)
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "string_value",
			FStrings: []string{"foo", "bar"},
			FInt32:   int32(i),
			FInt32S:  []int32{int32(i), int32(i + 1)},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub_string"},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   int64(i * 1000),
			FInt64S:  []int64{int64(i * 1000), int64(i * 2000)},
			FBytes:   []byte("byte_payload"),
			FBytess:  [][]byte{[]byte("a"), []byte("b")},
			FFloat:   float32(i) + 0.25,
			FFloats:  []float32{1.1, 2.2},
		}
		messages[i] = msg

		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	lastMsg := messages[len(messages)-1]

	if reply.FString != lastMsg.FString {
		t.Errorf("FString: got %v, want %v", reply.FString, lastMsg.FString)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("FInt32: got %v, want %v", reply.FInt32, lastMsg.FInt32)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("FEnum: got %v, want %v", reply.FEnum, lastMsg.FEnum)
	}
	if reply.FSub.GetFString() != lastMsg.FSub.GetFString() {
		t.Errorf("FSub.FString: got %v, want %v", reply.FSub.GetFString(), lastMsg.FSub.GetFString())
	}
	if reply.FBool != lastMsg.FBool {
		t.Errorf("FBool: got %v, want %v", reply.FBool, lastMsg.FBool)
	}
	if reply.FInt64 != lastMsg.FInt64 {
		t.Errorf("FInt64: got %v, want %v", reply.FInt64, lastMsg.FInt64)
	}
	if string(reply.FBytes) != string(lastMsg.FBytes) {
		t.Errorf("FBytes: got %v, want %v", reply.FBytes, lastMsg.FBytes)
	}
	if reply.FFloat != lastMsg.FFloat {
		t.Errorf("FFloat: got %v, want %v", reply.FFloat, lastMsg.FFloat)
	}

	// Additional checks for repeated fields can be added as needed
}
