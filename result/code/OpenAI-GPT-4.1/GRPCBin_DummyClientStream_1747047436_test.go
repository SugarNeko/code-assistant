package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
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

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	var lastMsg *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"foo", "bar"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_2,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub-string",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:   true,
			FBools:  []bool{true, false, true},
			FInt64:  int64(i * 10),
			FInt64S: []int64{10, 20, 30},
			FBytes:  []byte{0x61, 0x62, 0x63},
			FBytess: [][]byte{{0xDE, 0xAD}, {0xBE, 0xEF}},
			FFloat:  3.14,
			FFloats: []float32{1.5, 2.5},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
		lastMsg = msg
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive reply from stream: %v", err)
	}

	// Validate that the response equals the last sent message (server echoes last message)
	if reply.FString != lastMsg.FString {
		t.Errorf("Expected FString: %q, got: %q", lastMsg.FString, reply.FString)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("Expected FInt32: %d, got: %d", lastMsg.FInt32, reply.FInt32)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("Expected FEnum: %v, got: %v", lastMsg.FEnum, reply.FEnum)
	}
	// Add further field validations as needed
}
