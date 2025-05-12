package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream_Positive(t *testing.T) {
	// Connect to the gRPC server with 15s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("DummyClientStream create stream error: %v", err)
	}

	// Prepare and send 10 DummyMessage requests
	var lastMsg *grpcbin.DummyMessage
	for i := 1; i <= 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "message_" + string(rune(i)),
			FStrings:  []string{"foo", "bar"},
			FInt32:    int32(i),
			FInt32S:   []int32{int32(i), int32(i + 1)},
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub_" + string(rune(i))},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     (i%2 == 0),
			FBools:    []bool{true, false},
			FInt64:    int64(i * 1000),
			FInt64S:   []int64{int64(i), int64(i + 2)},
			FBytes:    []byte{1, 2, 3},
			FBytess:   [][]byte{{4, 5}, {6}},
			FFloat:    float32(i) * 1.1,
			FFloats:   []float32{2.3, 4.5},
		}
		lastMsg = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message %d: %v", i, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive reply: %v", err)
	}

	// Validate the reply matches the last sent message
	if reply.FString != lastMsg.FString {
		t.Errorf("FString mismatch: got %q, want %q", reply.FString, lastMsg.FString)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", reply.FInt32, lastMsg.FInt32)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", reply.FEnum, lastMsg.FEnum)
	}
	if reply.FSub == nil || lastMsg.FSub == nil || reply.FSub.FString != lastMsg.FSub.FString {
		t.Errorf("FSub FString mismatch: got %v, want %v", reply.FSub, lastMsg.FSub)
	}
	// More field validations can be added as needed
}
