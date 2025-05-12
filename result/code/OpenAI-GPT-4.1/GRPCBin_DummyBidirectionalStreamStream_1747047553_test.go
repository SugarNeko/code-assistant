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

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("DummyBidirectionalStreamStream open failed: %v", err)
	}

	testMsg := &pb.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"str1", "str2"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "sub1"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub2"}, {FString: "sub3"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    100,
		FInt64S:   []int64{101, 102},
		FBytes:    []byte{0x01, 0x02},
		FBytess:   [][]byte{{0x03}, {0x04, 0x05}},
		FFloat:    1.23,
		FFloats:   []float32{4.56, 7.89},
	}

	if err := stream.Send(testMsg); err != nil {
		t.Fatalf("Failed to send: %v", err)
	}

	recvMsg, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive: %v", err)
	}

	// Validate response fields match the sent fields
	if recvMsg.FString != testMsg.FString {
		t.Errorf("FString: got %v, want %v", recvMsg.FString, testMsg.FString)
	}
	if len(recvMsg.FStrings) != len(testMsg.FStrings) {
		t.Errorf("FStrings: got %v, want %v", recvMsg.FStrings, testMsg.FStrings)
	}
	if recvMsg.FInt32 != testMsg.FInt32 {
		t.Errorf("FInt32: got %v, want %v", recvMsg.FInt32, testMsg.FInt32)
	}
	if len(recvMsg.FInt32S) != len(testMsg.FInt32S) {
		t.Errorf("FInt32S: got %v, want %v", recvMsg.FInt32S, testMsg.FInt32S)
	}
	if recvMsg.FEnum != testMsg.FEnum {
		t.Errorf("FEnum: got %v, want %v", recvMsg.FEnum, testMsg.FEnum)
	}
	if len(recvMsg.FEnums) != len(testMsg.FEnums) {
		t.Errorf("FEnums: got %v, want %v", recvMsg.FEnums, testMsg.FEnums)
	}
	if (recvMsg.FSub == nil) != (testMsg.FSub == nil) ||
		(recvMsg.FSub != nil && recvMsg.FSub.FString != testMsg.FSub.FString) {
		t.Errorf("FSub: got %v, want %v", recvMsg.FSub, testMsg.FSub)
	}
	if len(recvMsg.FSubs) != len(testMsg.FSubs) {
		t.Errorf("FSubs: got %v, want %v", recvMsg.FSubs, testMsg.FSubs)
	}
	if recvMsg.FBool != testMsg.FBool {
		t.Errorf("FBool: got %v, want %v", recvMsg.FBool, testMsg.FBool)
	}
	if len(recvMsg.FBools) != len(testMsg.FBools) {
		t.Errorf("FBools: got %v, want %v", recvMsg.FBools, testMsg.FBools)
	}
	if recvMsg.FInt64 != testMsg.FInt64 {
		t.Errorf("FInt64: got %v, want %v", recvMsg.FInt64, testMsg.FInt64)
	}
	if len(recvMsg.FInt64S) != len(testMsg.FInt64S) {
		t.Errorf("FInt64S: got %v, want %v", recvMsg.FInt64S, testMsg.FInt64S)
	}
	if string(recvMsg.FBytes) != string(testMsg.FBytes) {
		t.Errorf("FBytes: got %v, want %v", recvMsg.FBytes, testMsg.FBytes)
	}
	if len(recvMsg.FBytess) != len(testMsg.FBytess) {
		t.Errorf("FBytess: got %v, want %v", recvMsg.FBytess, testMsg.FBytess)
	}
	if recvMsg.FFloat != testMsg.FFloat {
		t.Errorf("FFloat: got %v, want %v", recvMsg.FFloat, testMsg.FFloat)
	}
	if len(recvMsg.FFloats) != len(testMsg.FFloats) {
		t.Errorf("FFloats: got %v, want %v", recvMsg.FFloats, testMsg.FFloats)
	}

	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to CloseSend: %v", err)
	}
}
