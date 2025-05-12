package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyClientStream_Positive(t *testing.T) {
	addr := "grpcb.in:9000"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("failed to create client stream: %v", err)
	}

	var lastMsg *grpcbin.DummyMessage

	for i := 1; i <= 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "test-string",
			FStrings: []string{"foo", "bar"},
			FInt32:   int32(i),
			FInt32S:  []int32{int32(i * 2), int32(i * 3)},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2, grpcbin.DummyMessage_ENUM_0},
			FSub: &grpcbin.DummyMessage_Sub{
				FString: "sub-value",
			},
			FSubs: []*grpcbin.DummyMessage_Sub{
				{FString: "sub-a"},
				{FString: "sub-b"},
			},
			FBool:   i%2 == 0,
			FBools:  []bool{true, false, true},
			FInt64:  int64(i) * 100,
			FInt64S: []int64{int64(i) * 200, int64(i) * 300},
			FBytes:  []byte{0x01, 0x02, byte(i)},
			FBytess: [][]byte{{0x03, 0x04}, {0x05, byte(i)}},
			FFloat:  float32(i) + 0.5,
			FFloats: []float32{float32(i) * 1.1, float32(i) * 1.2},
		}
		lastMsg = msg
		if err := stream.Send(msg); err != nil {
			t.Fatalf("send failed at message %d: %v", i, err)
		}
	}
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("failed to close send side of the stream: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive from DummyClientStream: %v", err)
	}

	// Response validation: check that the response matches the last sent message
	if resp.FString != lastMsg.FString {
		t.Errorf("response FString = %s; want %s", resp.FString, lastMsg.FString)
	}
	if resp.FInt32 != lastMsg.FInt32 {
		t.Errorf("response FInt32 = %d; want %d", resp.FInt32, lastMsg.FInt32)
	}
	if resp.FEnum != lastMsg.FEnum {
		t.Errorf("response FEnum = %v; want %v", resp.FEnum, lastMsg.FEnum)
	}
	if resp.FSub == nil || lastMsg.FSub == nil || resp.FSub.FString != lastMsg.FSub.FString {
		t.Errorf("response FSub.FString = %v; want %v", resp.FSub, lastMsg.FSub)
	}
	if resp.FBool != lastMsg.FBool {
		t.Errorf("response FBool = %v; want %v", resp.FBool, lastMsg.FBool)
	}
	if resp.FInt64 != lastMsg.FInt64 {
		t.Errorf("response FInt64 = %d; want %d", resp.FInt64, lastMsg.FInt64)
	}
	if string(resp.FBytes) != string(lastMsg.FBytes) {
		t.Errorf("response FBytes = %v; want %v", resp.FBytes, lastMsg.FBytes)
	}
	if resp.FFloat != lastMsg.FFloat {
		t.Errorf("response FFloat = %v; want %v", resp.FFloat, lastMsg.FFloat)
	}
}
