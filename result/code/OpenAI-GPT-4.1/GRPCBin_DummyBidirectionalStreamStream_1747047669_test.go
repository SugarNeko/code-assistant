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
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("failed to create stream: %v", err)
	}

	req := &pb.DummyMessage{
		FString:  "hello",
		FStrings: []string{"foo", "bar"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:     &pb.DummyMessage_Sub{FString: "submessage"},
		FSubs:    []*pb.DummyMessage_Sub{{FString: "s1"}, {FString: "s2"}},
		FBool:    true,
		FBools:   []bool{true, false, true},
		FInt64:   123456789,
		FInt64S:  []int64{111, 222},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{[]byte("a"), []byte("b")},
		FFloat:   1.23,
		FFloats:  []float32{4.56, 7.89},
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("failed to send: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString = %q; want %q", resp.FString, req.FString)
	}
	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings len = %d; want %d", len(resp.FStrings), len(req.FStrings))
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 = %d; want %d", resp.FInt32, req.FInt32)
	}
	if len(resp.FInt32S) != len(req.FInt32S) {
		t.Errorf("FInt32S len = %d; want %d", len(resp.FInt32S), len(req.FInt32S))
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum = %v; want %v", resp.FEnum, req.FEnum)
	}
	if len(resp.FEnums) != len(req.FEnums) {
		t.Errorf("FEnums len = %d; want %d", len(resp.FEnums), len(req.FEnums))
	}
	if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
		t.Errorf("FSub = %+v; want %+v", resp.FSub, req.FSub)
	}
	if len(resp.FSubs) != len(req.FSubs) {
		t.Errorf("FSubs len = %d; want %d", len(resp.FSubs), len(req.FSubs))
	}
	if resp.FBool != req.FBool {
		t.Errorf("FBool = %v; want %v", resp.FBool, req.FBool)
	}
	if len(resp.FBools) != len(req.FBools) {
		t.Errorf("FBools len = %d; want %d", len(resp.FBools), len(req.FBools))
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("FInt64 = %d; want %d", resp.FInt64, req.FInt64)
	}
	if len(resp.FInt64S) != len(req.FInt64S) {
		t.Errorf("FInt64S len = %d; want %d", len(resp.FInt64S), len(req.FInt64S))
	}
	if string(resp.FBytes) != string(req.FBytes) {
		t.Errorf("FBytes = %v; want %v", resp.FBytes, req.FBytes)
	}
	if len(resp.FBytess) != len(req.FBytess) {
		t.Errorf("FBytess len = %d; want %d", len(resp.FBytess), len(req.FBytess))
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("FFloat = %v; want %v", resp.FFloat, req.FFloat)
	}
	if len(resp.FFloats) != len(req.FFloats) {
		t.Errorf("FFloats len = %d; want %d", len(resp.FFloats), len(req.FFloats))
	}
}
