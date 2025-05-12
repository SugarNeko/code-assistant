package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_1, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "nested"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   1001,
		FInt64S:  []int64{11, 12, 13},
		FBytes:   []byte("hello"),
		FBytess:  [][]byte{[]byte("a"), []byte("b")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send DummyMessage: %v", err)
	}

	rep, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive DummyMessage: %v", err)
	}

	// Validate echoed response matches sent message
	if rep.FString != req.FString {
		t.Errorf("FString mismatch: got %q, want %q", rep.FString, req.FString)
	}
	if len(rep.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(rep.FStrings), len(req.FStrings))
	}
	if rep.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", rep.FInt32, req.FInt32)
	}
	if len(rep.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %d, want %d", len(rep.FInt32S), len(req.FInt32S))
	}
	if rep.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", rep.FEnum, req.FEnum)
	}
	if len(rep.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(rep.FEnums), len(req.FEnums))
	}
	if rep.FSub == nil || req.FSub == nil || rep.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %q, want %q", rep.FSub.GetFString(), req.FSub.GetFString())
	}
	if len(rep.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(rep.FSubs), len(req.FSubs))
	}
	if rep.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", rep.FBool, req.FBool)
	}
	if len(rep.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(rep.FBools), len(req.FBools))
	}
	if rep.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", rep.FInt64, req.FInt64)
	}
	if len(rep.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %d, want %d", len(rep.FInt64S), len(req.FInt64S))
	}
	if string(rep.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %q, want %q", rep.FBytes, req.FBytes)
	}
	if len(rep.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(rep.FBytess), len(req.FBytess))
	}
	if rep.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", rep.FFloat, req.FFloat)
	}
	if len(rep.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(rep.FFloats), len(req.FFloats))
	}

	// Close stream send direction (not strictly necessary in bi-directional testing above)
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("CloseSend failed: %v", err)
	}
}
