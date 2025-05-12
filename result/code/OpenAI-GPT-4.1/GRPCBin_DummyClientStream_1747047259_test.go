package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	// Setup connect options with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect to grpcbin:9000: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("failed to open DummyClientStream: %v", err)
	}

	// Prepare 10 DummyMessage requests
	var lastSent *grpcbin.DummyMessage
	for i := 1; i <= 10; i++ {
		lastSent = &grpcbin.DummyMessage{
			FString:   "test-string",
			FStrings:  []string{"one", "two", "three"},
			FInt32:    int32(i),
			FInt32S:   []int32{1, 2, 3},
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-string"},
			FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub-0"}, {FString: "sub-1"}},
			FBool:     i%2 == 0,
			FBools:    []bool{true, false},
			FInt64:    int64(i * 10000),
			FInt64S:   []int64{123456789, 987654321},
			FBytes:    []byte("byte-data"),
			FBytess:   [][]byte{[]byte("bytes-1"), []byte("bytes-2")},
			FFloat:    float32(i) * 3.14,
			FFloats:   []float32{3.1, 4.2, 5.3},
		}
		if err := stream.Send(lastSent); err != nil {
			t.Fatalf("failed to send DummyMessage #%d: %v", i, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to CloseAndRecv: %v", err)
	}

	// Validate response matches last sent
	if reply == nil {
		t.Fatalf("Expected non-nil DummyMessage reply")
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

	// Optionally, validate more repeated/complex fields in a similar manner
}
