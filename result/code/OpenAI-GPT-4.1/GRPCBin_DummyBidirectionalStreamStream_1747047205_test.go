package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyBidirectionalStreamStream_Positive(t *testing.T) {
	addr := "grpcb.in:9000"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Prepare a typical DummyMessage as positive case
	req := &pb.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"foo", "bar"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_1},
		FSub:      &pb.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    987654321,
		FInt64S:   []int64{111, 222},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("one"), []byte("two")},
		FFloat:    3.14,
		FFloats:   []float32{2.71, 1.41},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send DummyMessage: %v", err)
	}
	// Optionally send another message
	// if err := stream.Send(req2); err != nil { ... }

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive DummyMessage: %v", err)
	}

	// Validate the response matches what was sent
	if resp.FString != req.FString {
		t.Errorf("response FString = %q, want %q", resp.FString, req.FString)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("response FInt32 = %d, want %d", resp.FInt32, req.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("response FEnum = %v, want %v", resp.FEnum, req.FEnum)
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("response FSub = %+v, want %+v", resp.FSub, req.FSub)
	}
	if resp.FBool != req.FBool {
		t.Errorf("response FBool = %v, want %v", resp.FBool, req.FBool)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("response FInt64 = %d, want %d", resp.FInt64, req.FInt64)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("response FBytes = %v, want %v", resp.FBytes, req.FBytes)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("response FFloat = %v, want %v", resp.FFloat, req.FFloat)
	}
	// Validate repeated and other fields as needed, e.g. len(), equality, etc.

	// Optionally, test stream can send and receive multiple times

	// Close the send direction of the stream
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to close send stream: %v", err)
	}
}
