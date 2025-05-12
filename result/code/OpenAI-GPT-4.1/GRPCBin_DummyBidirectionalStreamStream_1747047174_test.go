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
		t.Fatalf("failed to open stream: %v", err)
	}

	sendMsg := &grpcbin.DummyMessage{
		FString: "hello",
		FStrings: []string{"foo", "bar"},
		FInt32: 123,
		FInt32S: []int32{1, 2},
		FEnum: grpcbin.DummyMessage_ENUM_2,
		FEnums: []grpcbin.DummyMessage_Enum{
			grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_1,
		},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "sub-string",
		},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{FString: "sub1"},
			{FString: "sub2"},
		},
		FBool: true,
		FBools: []bool{true, false},
		FInt64: 1234567890,
		FInt64S: []int64{9, 8, 7},
		FBytes: []byte("bytes-data"),
		FBytess: [][]byte{
			[]byte("bytes1"), []byte("bytes2"),
		},
		FFloat: 3.14,
		FFloats: []float32{2.71, 1.41},
	}

	// Send the message to the stream
	if err := stream.Send(sendMsg); err != nil {
		t.Fatalf("failed to send message: %v", err)
	}

	// Receive echoed message
	recvMsg, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive message: %v", err)
	}

	// Validate response matches the sent message
	if recvMsg.FString != sendMsg.FString {
		t.Errorf("got FString=%v, want %v", recvMsg.FString, sendMsg.FString)
	}
	if len(recvMsg.FStrings) != len(sendMsg.FStrings) {
		t.Errorf("got FStrings len=%v, want %v", len(recvMsg.FStrings), len(sendMsg.FStrings))
	}
	if recvMsg.FInt32 != sendMsg.FInt32 {
		t.Errorf("got FInt32=%v, want %v", recvMsg.FInt32, sendMsg.FInt32)
	}
	if len(recvMsg.FInt32S) != len(sendMsg.FInt32S) {
		t.Errorf("got FInt32S len=%v, want %v", len(recvMsg.FInt32S), len(sendMsg.FInt32S))
	}
	if recvMsg.FEnum != sendMsg.FEnum {
		t.Errorf("got FEnum=%v, want %v", recvMsg.FEnum, sendMsg.FEnum)
	}
	if len(recvMsg.FEnums) != len(sendMsg.FEnums) {
		t.Errorf("got FEnums len=%v, want %v", len(recvMsg.FEnums), len(sendMsg.FEnums))
	}
	if recvMsg.FSub == nil || recvMsg.FSub.FString != sendMsg.FSub.FString {
		t.Errorf("got FSub=%v, want %v", recvMsg.FSub, sendMsg.FSub)
	}
	if len(recvMsg.FSubs) != len(sendMsg.FSubs) {
		t.Errorf("got FSubs len=%v, want %v", len(recvMsg.FSubs), len(sendMsg.FSubs))
	}
	if recvMsg.FBool != sendMsg.FBool {
		t.Errorf("got FBool=%v, want %v", recvMsg.FBool, sendMsg.FBool)
	}
	if len(recvMsg.FBools) != len(sendMsg.FBools) {
		t.Errorf("got FBools len=%v, want %v", len(recvMsg.FBools), len(sendMsg.FBools))
	}
	if recvMsg.FInt64 != sendMsg.FInt64 {
		t.Errorf("got FInt64=%v, want %v", recvMsg.FInt64, sendMsg.FInt64)
	}
	if len(recvMsg.FInt64S) != len(sendMsg.FInt64S) {
		t.Errorf("got FInt64S len=%v, want %v", len(recvMsg.FInt64S), len(sendMsg.FInt64S))
	}
	if string(recvMsg.FBytes) != string(sendMsg.FBytes) {
		t.Errorf("got FBytes=%v, want %v", recvMsg.FBytes, sendMsg.FBytes)
	}
	if len(recvMsg.FBytess) != len(sendMsg.FBytess) {
		t.Errorf("got FBytess len=%v, want %v", len(recvMsg.FBytess), len(sendMsg.FBytess))
	}
	if recvMsg.FFloat != sendMsg.FFloat {
		t.Errorf("got FFloat=%v, want %v", recvMsg.FFloat, sendMsg.FFloat)
	}
	if len(recvMsg.FFloats) != len(sendMsg.FFloats) {
		t.Errorf("got FFloats len=%v, want %v", len(recvMsg.FFloats), len(sendMsg.FFloats))
	}

	// clean up stream and context
	stream.CloseSend()
}
