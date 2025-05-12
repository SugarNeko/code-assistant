package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	// Set up a connection to the server with a 15-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("failed to open DummyClientStream: %v", err)
	}

	// Create 10 different DummyMessage requests
	var lastMessage *grpcbin.DummyMessage
	for i := 1; i <= 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "message " + string(rune(i)),
			FStrings:  []string{"foo", "bar"},
			FInt32:    int32(i),
			FInt32S:   []int32{int32(i * 10), int32(i * 20)},
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub_" + string(rune(i))},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     i%2 == 0,
			FBools:    []bool{true, false},
			FInt64:    int64(i * 1000),
			FInt64S:   []int64{int64(i * 100), int64(i * 200)},
			FBytes:    []byte("bytes"),
			FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
			FFloat:    3.14 * float32(i),
			FFloats:   []float32{1.1 * float32(i), 2.2 * float32(i)},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message %d: %v", i, err)
		}
		lastMessage = msg
	}

	// Close sending and receive response (should be last sent message)
	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive stream response: %v", err)
	}

	// Validate that the response matches the last message sent
	if resp.FString != lastMessage.FString {
		t.Errorf("unexpected FString: got %q, want %q", resp.FString, lastMessage.FString)
	}
	if resp.FInt32 != lastMessage.FInt32 {
		t.Errorf("unexpected FInt32: got %d, want %d", resp.FInt32, lastMessage.FInt32)
	}
	if resp.FEnum != lastMessage.FEnum {
		t.Errorf("unexpected FEnum: got %v, want %v", resp.FEnum, lastMessage.FEnum)
	}
	if resp.FSub.GetFString() != lastMessage.FSub.GetFString() {
		t.Errorf("unexpected FSub.FString: got %q, want %q", resp.FSub.GetFString(), lastMessage.FSub.GetFString())
	}
	if resp.FBool != lastMessage.FBool {
		t.Errorf("unexpected FBool: got %v, want %v", resp.FBool, lastMessage.FBool)
	}
	if resp.FInt64 != lastMessage.FInt64 {
		t.Errorf("unexpected FInt64: got %d, want %d", resp.FInt64, lastMessage.FInt64)
	}
	if string(resp.FBytes) != string(lastMessage.FBytes) {
		t.Errorf("unexpected FBytes: got %v, want %v", resp.FBytes, lastMessage.FBytes)
	}
	if resp.FFloat != lastMessage.FFloat {
		t.Errorf("unexpected FFloat: got %v, want %v", resp.FFloat, lastMessage.FFloat)
	}

	// (Add more field checks as needed for thorough validation)
}
