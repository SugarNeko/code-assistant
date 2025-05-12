package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to open stream: %v", err)
	}

	var lastMsg *pb.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"a", "b", "c"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    pb.DummyMessage_ENUM_2,
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_1},
			FSub:     &pb.DummyMessage_Sub{FString: "substring"},
			FSubs:    []*pb.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
			FBool:    i%2 == 0,
			FBools:   []bool{true, false, true},
			FInt64:   int64(i * 100),
			FInt64S:  []int64{10, 20, 30},
			FBytes:   []byte{0x1, 0x2},
			FBytess:  [][]byte{{0x3, 0x4}, {0x5, 0x6}},
			FFloat:   float32(i) + 0.5,
			FFloats:  []float32{1.5, 2.5, 3.5},
		}
		lastMsg = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i+1, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Error on CloseAndRecv: %v", err)
	}

	// Client response validation: Check that the reply matches the last sent message
	if reply.GetFString() != lastMsg.GetFString() ||
		reply.GetFInt32() != lastMsg.GetFInt32() ||
		reply.GetFEnum() != lastMsg.GetFEnum() ||
		reply.GetFBool() != lastMsg.GetFBool() ||
		reply.GetFInt64() != lastMsg.GetFInt64() ||
		reply.GetFFloat() != lastMsg.GetFFloat() {
		t.Errorf("Server reply does not match last sent message.\nGot: %+v\nWant: %+v", reply, lastMsg)
	}

	// Check repeated and nested fields
	if len(reply.GetFStrings()) != len(lastMsg.GetFStrings()) {
		t.Errorf("Mismatch in FStrings: got %d, want %d", len(reply.GetFStrings()), len(lastMsg.GetFStrings()))
	}
	if reply.GetFSub().GetFString() != lastMsg.GetFSub().GetFString() {
		t.Errorf("Mismatch in FSub.FString: got %s, want %s", reply.GetFSub().GetFString(), lastMsg.GetFSub().GetFString())
	}
}
