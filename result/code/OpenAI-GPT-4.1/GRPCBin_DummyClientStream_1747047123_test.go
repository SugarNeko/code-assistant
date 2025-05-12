package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	address := "grpcb.in:9000"

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to open DummyClientStream: %v", err)
	}

	var lastMsg *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "string_value",
			FStrings: []string{"s1", "s2"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub_string",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub1"},
				{FString: "sub2"},
			},
			FBool:   (i%2 == 0),
			FBools:  []bool{true, false},
			FInt64:  int64(i),
			FInt64S: []int64{100, 200},
			FBytes:  []byte{0x10, 0x20},
			FBytess: [][]byte{{0x01, 0x02}, {0x03, 0x04}},
			FFloat:  float32(1.23),
			FFloats: []float32{2.34, 3.45},
		}
		lastMsg = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive DummyClientStream response: %v", err)
	}

	// Client Response Validation (simple, limited validation)
	if reply == nil {
		t.Fatalf("Expected non-nil reply from server")
	}
	if reply.FString != lastMsg.FString {
		t.Errorf("Expected last message FString=%q, got %q", lastMsg.FString, reply.FString)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("Expected last message FInt32=%d, got %d", lastMsg.FInt32, reply.FInt32)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("Expected last message FEnum=%v, got %v", lastMsg.FEnum, reply.FEnum)
	}
	if reply.FBool != lastMsg.FBool {
		t.Errorf("Expected last message FBool=%v, got %v", lastMsg.FBool, reply.FBool)
	}
}
