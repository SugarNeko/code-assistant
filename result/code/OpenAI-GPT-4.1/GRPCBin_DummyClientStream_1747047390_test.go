package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	// Connection with 15s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create client stream: %v", err)
	}

	// Prepare 10 DummyMessages
	var lastMsg *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"a", "b"},
			FInt32:   int32(i),
			FInt32S:  []int32{int32(i), int32(i + 1)},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-field"},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    i%2 == 0,
			FBools:   []bool{i%2 == 0, i%2 != 0},
			FInt64:   int64(i * 100),
			FInt64S:  []int64{int64(i * 10), int64(i * 100)},
			FBytes:   []byte("bytes"),
			FBytess:  [][]byte{[]byte("a"), []byte("b")},
			FFloat:   float32(i) * 1.1,
			FFloats:  []float32{0.2, 1.2, 2.2},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message #%d: %v", i, err)
		}
		lastMsg = msg
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response from server: %v", err)
	}

	// Validate that reply equals last message sent
	if reply.FString != lastMsg.FString {
		t.Errorf("Expected FString: %s, got: %s", lastMsg.FString, reply.FString)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("Expected FInt32: %d, got: %d", lastMsg.FInt32, reply.FInt32)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("Expected FEnum: %v, got: %v", lastMsg.FEnum, reply.FEnum)
	}
	if reply.FBool != lastMsg.FBool {
		t.Errorf("Expected FBool: %v, got: %v", lastMsg.FBool, reply.FBool)
	}
	if reply.FInt64 != lastMsg.FInt64 {
		t.Errorf("Expected FInt64: %v, got: %v", lastMsg.FInt64, reply.FInt64)
	}
	if reply.FFloat != lastMsg.FFloat {
		t.Errorf("Expected FFloat: %v, got: %v", lastMsg.FFloat, reply.FFloat)
	}
	// Additional field validations can be added as needed
}
