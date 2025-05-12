package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("Failed to open stream: %v", err)
	}

	// Build a typical DummyMessage fully populated
	req := &pb.DummyMessage{
		FString:   "test string",
		FStrings:  []string{"str1", "str2"},
		FInt32:    123,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
		FSub:      &pb.DummyMessage_Sub{FString: "sub str"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    1234567890,
		FInt64S:   []int64{10, 20, 30},
		FBytes:    []byte("bytes-data"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    1.234,
		FFloats:   []float32{2.345, 3.456},
	}

	// Send the message
	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send DummyMessage: %v", err)
	}

	// Receive the echoed message from server
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive DummyMessage: %v", err)
	}

	// Client response validation
	if resp.FString != req.FString ||
		len(resp.FStrings) != len(req.FStrings) ||
		resp.FInt32 != req.FInt32 ||
		len(resp.FInt32S) != len(req.FInt32S) ||
		resp.FEnum != req.FEnum ||
		len(resp.FEnums) != len(req.FEnums) ||
		resp.FSub.GetFString() != req.FSub.GetFString() ||
		len(resp.FSubs) != len(req.FSubs) ||
		resp.FBool != req.FBool ||
		len(resp.FBools) != len(req.FBools) ||
		resp.FInt64 != req.FInt64 ||
		len(resp.FInt64S) != len(req.FInt64S) ||
		string(resp.FBytes) != string(req.FBytes) ||
		len(resp.FBytess) != len(req.FBytess) ||
		resp.FFloat != req.FFloat ||
		len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("Response does not match request; got %+v, want %+v", resp, req)
	}
}
