package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to create bidirectional stream: %v", err)
	}

	req := &pb.DummyMessage{
		FString:    "hello",
		FStrings:   []string{"foo", "bar"},
		FInt32:     123,
		FInt32S:    []int32{1, 2},
		FEnum:      pb.DummyMessage_ENUM_1,
		FEnums:     []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:       &pb.DummyMessage_Sub{FString: "subval"},
		FSubs:      []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:      true,
		FBools:     []bool{true, false},
		FInt64:     99,
		FInt64S:    []int64{100, 200},
		FBytes:     []byte("data"),
		FBytess:    [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:     3.14,
		FFloats:    []float32{1.2, 3.4},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	// Client response validation: check if response matches the sent request
	if got, want := resp.FString, req.FString; got != want {
		t.Errorf("FString: got %q, want %q", got, want)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings: got %v, want %v", resp.FStrings, req.FStrings)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32: got %d, want %d", resp.FInt32, req.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum: got %v, want %v", resp.FEnum, req.FEnum)
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub: got %+v, want %+v", resp.FSub, req.FSub)
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool: got %t, want %t", resp.FBool, req.FBool)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64: got %d, want %d", resp.FInt64, req.FInt64)
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes: got %q, want %q", resp.FBytes, req.FBytes)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat: got %f, want %f", resp.FFloat, req.FFloat)
	}

	// Add more detailed checks as per requirements...
	if err := stream.CloseSend(); err != nil {
		t.Errorf("Failed to close stream: %v", err)
	}
}
