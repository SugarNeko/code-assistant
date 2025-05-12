package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("failed to open stream: %v", err)
	}

	req := &pb.DummyMessage{
		FString:   "test-string",
		FStrings:  []string{"one", "two"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
		FSub:      &pb.DummyMessage_Sub{FString: "sub-string"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub-1"}, {FString: "sub-2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    987654321,
		FInt64S:   []int64{111, 222},
		FBytes:    []byte("byte-string"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2, 3.3},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	// Validate the response matches the request
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: got %s, want %s", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch: got %d, want %d", len(resp.FStrings), len(req.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch: got %d, want %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S length mismatch: got %d, want %d", len(resp.FInt32S), len(req.FInt32S))
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums length mismatch: got %d, want %d", len(resp.FEnums), len(req.FEnums))
	}
	if (resp.FSub == nil) != (req.FSub == nil) {
		t.Errorf("FSub nil mismatch: got %v, want %v", resp.FSub, req.FSub)
	}
	if req.FSub != nil && resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub.FString mismatch: got %s, want %s", resp.FSub.FString, req.FSub.FString)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs length mismatch: got %d, want %d", len(resp.FSubs), len(req.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools length mismatch: got %d, want %d", len(resp.FBools), len(req.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 mismatch: got %d, want %d", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S length mismatch: got %d, want %d", len(resp.FInt64S), len(req.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess length mismatch: got %d, want %d", len(resp.FBytess), len(req.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat mismatch: got %f, want %f", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats length mismatch: got %d, want %d", len(resp.FFloats), len(req.FFloats))
	}

	// Check stream close/EOF, and server closes cleanly
	if err := stream.CloseSend(); err != nil {
		t.Errorf("CloseSend failed: %v", err)
	}
	_, err = stream.Recv()
	if err == nil {
		t.Error("expected EOF from server")
	} else if _, ok := status.FromError(err); !ok && err.Error() != "EOF" {
		t.Errorf("unexpected stream close error: %v", err)
	}
}
