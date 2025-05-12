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
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"foo", "bar"},
		FInt32:   123,
		FInt32S:  []int32{4, 5},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   987654321,
		FInt64S:  []int64{1, 2, 3},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   3.14,
		FFloats:  []float32{2.71, 1.61},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("failed to send message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive message: %v", err)
	}

	// Validate response == request (since grpcbin echoes the message)
	if resp.FString != req.FString {
		t.Errorf("FString: got %q, want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings: got len %d, want %d", len(resp.FStrings), len(req.FStrings))
	}
	for i, v := range req.FStrings {
		if resp.FStrings[i] != v {
			t.Errorf("FStrings[%d]: got %q, want %q", i, resp.FStrings[i], v)
		}
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32: got %v, want %v", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S: got len %d, want %d", len(resp.FInt32S), len(req.FInt32S))
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums: got len %d, want %d", len(resp.FEnums), len(req.FEnums))
	}
	if resp.FSub == nil || req.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString: got %q, want %q", resp.FSub.FString, req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs: got len %d, want %d", len(resp.FSubs), len(req.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools: got len %d, want %d", len(resp.FBools), len(req.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64: got %v, want %v", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S: got len %d, want %d", len(resp.FInt64S), len(req.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes: got %q, want %q", string(resp.FBytes), string(req.FBytes))
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess: got len %d, want %d", len(resp.FBytess), len(req.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat: got %v, want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats: got len %d, want %d", len(resp.FFloats), len(req.FFloats))
	}
}
