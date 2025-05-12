package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to open stream: %v", err)
	}

	msg := &grpcbin.DummyMessage{
		FString:     "hello",
		FStrings:    []string{"foo", "bar"},
		FInt32:      42,
		FInt32S:     []int32{1, 2, 3},
		FEnum:       grpcbin.DummyMessage_ENUM_2,
		FEnums:      []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub:        &grpcbin.DummyMessage_Sub{FString: "sub-hello"},
		FSubs:       []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:       true,
		FBools:      []bool{true, false, true},
		FInt64:      123456789,
		FInt64S:     []int64{11, 22, 33},
		FBytes:      []byte("byte-data"),
		FBytess:     [][]byte{[]byte("foo"), []byte("bar")},
		FFloat:      42.5,
		FFloats:     []float32{1.23, 4.56},
	}

	if err := stream.Send(msg); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}
	recv, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}
	
	// Validate response matches sent message fields (Positive Case)
	if recv.FString != msg.FString {
		t.Errorf("FString mismatch: got %v want %v", recv.FString, msg.FString)
	}
	if len(recv.FStrings) != len(msg.FStrings) {
		t.Errorf("FStrings length mismatch: got %v want %v", len(recv.FStrings), len(msg.FStrings))
	}
	if recv.FInt32 != msg.FInt32 {
		t.Errorf("FInt32 mismatch: got %v want %v", recv.FInt32, msg.FInt32)
	}
	if recv.FEnum != msg.FEnum {
		t.Errorf("FEnum mismatch: got %v want %v", recv.FEnum, msg.FEnum)
	}
	if recv.FSub == nil || recv.FSub.FString != msg.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v want %v", recv.FSub, msg.FSub.FString)
	}
	if recv.FBool != msg.FBool {
		t.Errorf("FBool mismatch: got %v want %v", recv.FBool, msg.FBool)
	}
	if recv.FInt64 != msg.FInt64 {
		t.Errorf("FInt64 mismatch: got %v want %v", recv.FInt64, msg.FInt64)
	}
	if string(recv.FBytes) != string(msg.FBytes) {
		t.Errorf("FBytes mismatch: got %v want %v", string(recv.FBytes), string(msg.FBytes))
	}
	if recv.FFloat != msg.FFloat {
		t.Errorf("FFloat mismatch: got %v want %v", recv.FFloat, msg.FFloat)
	}

	// Close send direction
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to CloseSend: %v", err)
	}
}
