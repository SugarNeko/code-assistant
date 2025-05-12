package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect to grpc server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create DummyClientStream: %v", err)
	}

	expected := &grpcbin.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"hi", "world"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "submsg"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "a"}, {FString: "b"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    98765,
		FInt64S:   []int64{100, 200},
		FBytes:    []byte("bytestring"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	// send 9 dummy messages
	for i := 0; i < 9; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "dummy",
			FInt32:    int32(i),
			FEnum:     grpcbin.DummyMessage_ENUM_0,
			FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message %d: %v", i+1, err)
		}
	}
	// send the expected message last
	if err := stream.Send(expected); err != nil {
		t.Fatalf("Failed to send last expected message: %v", err)
	}
	// Close stream and receive response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response from stream: %v", err)
	}

	// Validate the response is equal to the last message sent (expected)
	if resp.FString != expected.FString {
		t.Errorf("FString mismatch: got %q, want %q", resp.FString, expected.FString)
	}
	if len(resp.FStrings) != len(expected.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(expected.FStrings))
	}
	if resp.FInt32 != expected.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, expected.FInt32)
	}
	if resp.FEnum != expected.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, expected.FEnum)
	}
	if resp.FSub == nil || resp.FSub.FString != expected.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %+v, want %+v", resp.FSub, expected.FSub)
	}
	if resp.FBool != expected.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, expected.FBool)
	}
	if resp.FInt64 != expected.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, expected.FInt64)
	}
	if string(resp.FBytes) != string(expected.FBytes) {
		t.Errorf("FBytes mismatch: got %s, want %s", string(resp.FBytes), string(expected.FBytes))
	}
	if resp.FFloat != expected.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, expected.FFloat)
	}
}
