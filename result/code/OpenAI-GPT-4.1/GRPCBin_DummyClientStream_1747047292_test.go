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
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("DummyClientStream open failed: %v", err)
	}

	var lastSent *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test_string",
			FStrings: []string{"str1", "str2"},
			FInt32:   int32(i),
			FInt32S:  []int32{int32(i), int32(i * 2)},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    i%2 == 0,
			FBools:   []bool{true, false},
			FInt64:   int64(i),
			FInt64S:  []int64{int64(i), int64(i * 3)},
			FBytes:   []byte("bytes"),
			FBytess:  [][]byte{[]byte("a"), []byte("b")},
			FFloat:   float32(i) + 0.1,
			FFloats:  []float32{2.3, 4.5},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message #%d: %v", i, err)
		}
		lastSent = msg
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive reply: %v", err)
	}

	if reply == nil {
		t.Fatal("reply is nil")
	}

	if reply.FString != lastSent.FString {
		t.Errorf("FString mismatch: got %q, want %q", reply.FString, lastSent.FString)
	}
	if reply.FInt32 != lastSent.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", reply.FInt32, lastSent.FInt32)
	}
	if reply.FEnum != lastSent.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", reply.FEnum, lastSent.FEnum)
	}
	// Add more detailed field comparisons as needed for further validation...
}
