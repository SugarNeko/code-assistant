package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("open stream error: %v", err)
	}

	sendMsg := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"world", "grpc"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "subfield"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{false, true},
		FInt64:    9876543210,
		FInt64S:   []int64{100, 200, 300},
		FBytes:    []byte("foo"),
		FBytess:   [][]byte{[]byte("bar"), []byte("baz")},
		FFloat:    1.23,
		FFloats:   []float32{4.56, 7.89},
	}
	if err := stream.Send(sendMsg); err != nil {
		t.Fatalf("failed to send: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive: %v", err)
	}

	if resp.FString != sendMsg.FString {
		t.Errorf("FString not match, got %q want %q", resp.FString, sendMsg.FString)
	}
	if len(resp.FStrings) != len(sendMsg.FStrings) {
		t.Errorf("FStrings length not match, got %d want %d", len(resp.FStrings), len(sendMsg.FStrings))
	}
	if resp.FInt32 != sendMsg.FInt32 {
		t.Errorf("FInt32 not match, got %d want %d", resp.FInt32, sendMsg.FInt32)
	}
	if len(resp.FInt32S) != len(sendMsg.FInt32S) {
		t.Errorf("FInt32S length not match, got %d want %d", len(resp.FInt32S), len(sendMsg.FInt32S))
	}
	if resp.FEnum != sendMsg.FEnum {
		t.Errorf("FEnum not match, got %v want %v", resp.FEnum, sendMsg.FEnum)
	}
	if len(resp.FEnums) != len(sendMsg.FEnums) {
		t.Errorf("FEnums length not match, got %d want %d", len(resp.FEnums), len(sendMsg.FEnums))
	}
	if resp.FSub == nil || resp.FSub.FString != sendMsg.FSub.FString {
		t.Errorf("FSub.FString not match, got %q want %q", resp.FSub.GetFString(), sendMsg.FSub.FString)
	}
	if len(resp.FSubs) != len(sendMsg.FSubs) {
		t.Errorf("FSubs length not match, got %d want %d", len(resp.FSubs), len(sendMsg.FSubs))
	}
	if resp.FBool != sendMsg.FBool {
		t.Errorf("FBool not match, got %v want %v", resp.FBool, sendMsg.FBool)
	}
	if len(resp.FBools) != len(sendMsg.FBools) {
		t.Errorf("FBools length not match, got %d want %d", len(resp.FBools), len(sendMsg.FBools))
	}
	if resp.FInt64 != sendMsg.FInt64 {
		t.Errorf("FInt64 not match, got %d want %d", resp.FInt64, sendMsg.FInt64)
	}
	if len(resp.FInt64S) != len(sendMsg.FInt64S) {
		t.Errorf("FInt64S length not match, got %d want %d", len(resp.FInt64S), len(sendMsg.FInt64S))
	}
	if string(resp.FBytes) != string(sendMsg.FBytes) {
		t.Errorf("FBytes not match, got %q want %q", resp.FBytes, sendMsg.FBytes)
	}
	if len(resp.FBytess) != len(sendMsg.FBytess) {
		t.Errorf("FBytess length not match, got %d want %d", len(resp.FBytess), len(sendMsg.FBytess))
	}
	if resp.FFloat != sendMsg.FFloat {
		t.Errorf("FFloat not match, got %v want %v", resp.FFloat, sendMsg.FFloat)
	}
	if len(resp.FFloats) != len(sendMsg.FFloats) {
		t.Errorf("FFloats length not match, got %d want %d", len(resp.FFloats), len(sendMsg.FFloats))
	}

	if err := stream.CloseSend(); err != nil {
		t.Errorf("failed to close stream: %v", err)
	}
}
