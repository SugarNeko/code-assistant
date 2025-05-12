package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyClientStream_Positive(t *testing.T) {
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
		t.Fatalf("DummyClientStream failed: %v", err)
	}

	var lastSent *grpcbin.DummyMessage

	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "hello",
			FStrings: []string{"a", "b"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "subfield"},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sublist1"}, {FString: "sublist2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   int64(i),
			FInt64S:  []int64{100, 200},
			FBytes:   []byte{0x01, 0x02},
			FBytess:  [][]byte{{0x03}, {0x04}},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2},
		}
		lastSent = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("failed to send message %d: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	// Client response validation: verify the response matches the last sent message
	if reply.FString != lastSent.FString {
		t.Errorf("expected FString %q, got %q", lastSent.FString, reply.FString)
	}
	if reply.FInt32 != lastSent.FInt32 {
		t.Errorf("expected FInt32 %v, got %v", lastSent.FInt32, reply.FInt32)
	}
	if reply.FEnum != lastSent.FEnum {
		t.Errorf("expected FEnum %v, got %v", lastSent.FEnum, reply.FEnum)
	}
	if reply.FSub == nil || lastSent.FSub == nil || reply.FSub.FString != lastSent.FSub.FString {
		t.Errorf("expected FSub.Fprintf %q, got %q", lastSent.FSub.FString, reply.FSub.FString)
	}
	if reply.FBool != lastSent.FBool {
		t.Errorf("expected FBool %v, got %v", lastSent.FBool, reply.FBool)
	}
	if reply.FInt64 != lastSent.FInt64 {
		t.Errorf("expected FInt64 %v, got %v", lastSent.FInt64, reply.FInt64)
	}
	if len(reply.FBytes) != len(lastSent.FBytes) {
		t.Errorf("expected FBytes len %d, got %d", len(lastSent.FBytes), len(reply.FBytes))
	}
	if reply.FFloat != lastSent.FFloat {
		t.Errorf("expected FFloat %v, got %v", lastSent.FFloat, reply.FFloat)
	}

	// (Optionally add further detailed comparisons for other fields as needed)
}
