package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyClientStream_Positive(t *testing.T) {
	// Connection with 15s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("DummyClientStream open stream failed: %v", err)
	}

	// Prepare 10 DummyMessages (compliant request)
	var lastMsg *pb.DummyMessage
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString:  "test message",
			FStrings: []string{"str1", "str2"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    pb.DummyMessage_ENUM_1,
			FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
			FSub:     &pb.DummyMessage_Sub{FString: "sub str"},
			FSubs:    []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    true,
			FBools:   []bool{true, false},
			FInt64:   int64(i * 1000),
			FInt64S:  []int64{123, 456},
			FBytes:   []byte{0xFF, 0xEE, 0xDD},
			FBytess:  [][]byte{{0x11, 0x22}, {0x33, 0x44}},
			FFloat:   float32(i) * 1.5,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}
		lastMsg = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Send failed on message %d: %v", i, err)
		}
	}
	// Close send direction, receive response
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("CloseAndRecv failed: %v", err)
	}

	// Validate response equals the last message sent (server echoes last)
	if reply == nil {
		t.Fatal("Received nil reply from DummyClientStream")
	}

	if reply.FString != lastMsg.FString {
		t.Errorf("Expected FString %q, got %q", lastMsg.FString, reply.FString)
	}
	if reply.FInt32 != lastMsg.FInt32 {
		t.Errorf("Expected FInt32 %d, got %d", lastMsg.FInt32, reply.FInt32)
	}
	if reply.FEnum != lastMsg.FEnum {
		t.Errorf("Expected FEnum %v, got %v", lastMsg.FEnum, reply.FEnum)
	}
	if reply.FBool != lastMsg.FBool {
		t.Errorf("Expected FBool %v, got %v", lastMsg.FBool, reply.FBool)
	}
	if reply.FInt64 != lastMsg.FInt64 {
		t.Errorf("Expected FInt64 %d, got %d", lastMsg.FInt64, reply.FInt64)
	}
	if reply.FFloat != lastMsg.FFloat {
		t.Errorf("Expected FFloat %v, got %v", lastMsg.FFloat, reply.FFloat)
	}
	// Add more validations as needed for other fields
}
