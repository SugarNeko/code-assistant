package grpcbin_test

import (
	"context"
	"testing"
	"time"
	"io"
	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to start stream: %v", err)
	}

	// Prepare a positive DummyMessage following proto schema
	req := &grpcbin.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"str1", "str2"},
		FInt32:    123,
		FInt32S:   []int32{100, 200},
		FEnum:     grpcbin.DummyMessage_ENUM_2,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    987654321,
		FInt64S:   []int64{111, 222},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	// Send request
	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive response and validate it's the same as sent
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %v, want %v", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %v, want %v", len(resp.FStrings), len(req.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %v, want %v", len(resp.FInt32S), len(req.FInt32S))
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %v, want %v", len(resp.FEnums), len(req.FEnums))
	}
	if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %v, want %v", resp.FSub.GetFString(), req.FSub.GetFString())
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %v, want %v", len(resp.FSubs), len(req.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %v, want %v", len(resp.FBools), len(req.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %v, want %v", len(resp.FInt64S), len(req.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", string(resp.FBytes), string(req.FBytes))
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %v, want %v", len(resp.FBytess), len(req.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %v, want %v", len(resp.FFloats), len(req.FFloats))
	}

	// Optionally test receive after closing send - should get EOF
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to close send: %v", err)
	}
	_, err = stream.Recv()
	if err != io.EOF {
		t.Errorf("Expected EOF after closing stream, got %v", err)
	}
}
