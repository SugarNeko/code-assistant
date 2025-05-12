package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to start DummyClientStream: %v", err)
	}

	var lastSent *grpcbin.DummyMessage
	for i := 1; i <= 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"foo", "bar"}, 
			FInt32:   int32(i),
			FInt32S:  []int32{int32(i), int32(i + 1)},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-string"},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    i%2 == 0,
			FBools:   []bool{i%2 == 0, i%2 != 0},
			FInt64:   int64(i * 10),
			FInt64S:  []int64{int64(i * 10), int64(i * 20)},
			FBytes:   []byte{0x01, byte(i)},
			FBytess:  [][]byte{{0x02, byte(i)}, {0x03, byte(i + 1)}},
			FFloat:   float32(i) * 3.14,
			FFloats:  []float32{float32(i), float32(i+1)},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
		lastSent = msg
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response from DummyClientStream: %v", err)
	}

	// Validate: server should respond with the last message sent
	if reply == nil {
		t.Fatal("Received nil response from server")
	}
	if reply.FString != lastSent.FString {
		t.Errorf("Expected FString %q, got %q", lastSent.FString, reply.FString)
	}
	if reply.FInt32 != lastSent.FInt32 {
		t.Errorf("Expected FInt32 %d, got %d", lastSent.FInt32, reply.FInt32)
	}
	if reply.FEnum != lastSent.FEnum {
		t.Errorf("Expected FEnum %v, got %v", lastSent.FEnum, reply.FEnum)
	}
	if reply.FSub == nil || lastSent.FSub == nil || reply.FSub.FString != lastSent.FSub.FString {
		t.Errorf("Expected FSub.FString %q, got %q", lastSent.FSub.FString, reply.FSub.FString)
	}
	if reply.FBool != lastSent.FBool {
		t.Errorf("Expected FBool %v, got %v", lastSent.FBool, reply.FBool)
	}
	if reply.FInt64 != lastSent.FInt64 {
		t.Errorf("Expected FInt64 %v, got %v", lastSent.FInt64, reply.FInt64)
	}
	if string(reply.FBytes) != string(lastSent.FBytes) {
		t.Errorf("Expected FBytes %v, got %v", lastSent.FBytes, reply.FBytes)
	}
	if reply.FFloat != lastSent.FFloat {
		t.Errorf("Expected FFloat %v, got %v", lastSent.FFloat, reply.FFloat)
	}

	// Add more thorough field validations as needed.
}
