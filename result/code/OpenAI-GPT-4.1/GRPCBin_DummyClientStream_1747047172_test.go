package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	// Setup connection with 15 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("failed to open client stream: %v", err)
	}

	// Prepare 10 DummyMessages according to the proto specification
	var lastMsg *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "f_string_value",
			FStrings:  []string{"s1", "s2"},
			FInt32:    int32(i),
			FInt32S:   []int32{1, 2, 3},
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_1},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub_string"},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     i%2 == 0,
			FBools:    []bool{true, false},
			FInt64:    int64(i * 100),
			FInt64S:   []int64{111, 222},
			FBytes:    []byte("bytes_here"),
			FBytess:   [][]byte{[]byte("a"), []byte("b")},
			FFloat:    float32(i) * 1.25,
			FFloats:   []float32{2.3, 4.5},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message #%d: %v", i, err)
		}
		lastMsg = msg
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to close stream and receive: %v", err)
	}

	// Response Validation
	if lastMsg == nil {
		t.Fatalf("lastMsg is nil")
	}
	if resp.FString != lastMsg.FString {
		t.Errorf("unexpected FString: got %q, want %q", resp.FString, lastMsg.FString)
	}
	if resp.FInt32 != lastMsg.FInt32 {
		t.Errorf("unexpected FInt32: got %d, want %d", resp.FInt32, lastMsg.FInt32)
	}
	if resp.FEnum != lastMsg.FEnum {
		t.Errorf("unexpected FEnum: got %v, want %v", resp.FEnum, lastMsg.FEnum)
	}
	// Add more field checks as required for thoroughness...
}
