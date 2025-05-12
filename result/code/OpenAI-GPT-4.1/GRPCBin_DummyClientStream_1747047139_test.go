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

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	var lastMsg *grpcbin.DummyMessage
	for i := 1; i <= 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "test-string",
			FStrings:  []string{"str1", "str2"},
			FInt32:    int32(i),
			FInt32S:   []int32{int32(i), int32(i + 1)},
			FEnum:     grpcbin.DummyMessage_ENUM_2,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-value"},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:     i%2 == 0,
			FBools:    []bool{true, false},
			FInt64:    int64(i * 10),
			FInt64S:   []int64{int64(i * 20), int64(i * 30)},
			FBytes:    []byte{0x01, 0x02},
			FBytess:   [][]byte{{0x03, 0x04}, {0x05, 0x06}},
			FFloat:    float32(i) * 1.44,
			FFloats:   []float32{3.14, 1.41},
		}
		lastMsg = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message #%d: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive final message: %v", err)
	}

	// Validate response is equal to the last sent message
	if reply.GetFString() != lastMsg.GetFString() {
		t.Errorf("f_string: got %v, want %v", reply.GetFString(), lastMsg.GetFString())
	}
	if reply.GetFInt32() != lastMsg.GetFInt32() {
		t.Errorf("f_int32: got %v, want %v", reply.GetFInt32(), lastMsg.GetFInt32())
	}
	if reply.GetFEnum() != lastMsg.GetFEnum() {
		t.Errorf("f_enum: got %v, want %v", reply.GetFEnum(), lastMsg.GetFEnum())
	}
	// ... (extend validation for other fields as needed)
}
