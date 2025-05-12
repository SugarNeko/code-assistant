package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestGRPCBin_DummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
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
			FInt32S:  []int32{int32(i), int32(i + 1)},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_1},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-str"},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"}, {FString: "sub2"},
			},
			FBool:    i%2 == 0,
			FBools:   []bool{true, false, true},
			FInt64:   int64(i + 100),
			FInt64S:  []int64{101, 102, 103},
			FBytes:   []byte{0x01, 0x02},
			FBytess:  [][]byte{{0x0a}, {0x0b}},
			FFloat:   1.23 + float32(i),
			FFloats:  []float32{2.34, 3.45, 4.56},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
		lastMsg = msg
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive reply: %v", err)
	}

	if reply == nil {
		t.Fatalf("Expected non-nil reply")
	}
	// Validate some client->server->client roundtrip fields
	if reply.FString != lastMsg.FString {
		t.Errorf("reply.FString = %q, want %q", reply.FString, lastMsg.FString)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("reply.FInt32 = %v, want %v", reply.FInt32, lastMsg.FInt32)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("reply.FEnum = %v, want %v", reply.FEnum, lastMsg.FEnum)
	}
	if reply.FBool != lastMsg.FBool {
		t.Errorf("reply.FBool = %v, want %v", reply.FBool, lastMsg.FBool)
	}
}
