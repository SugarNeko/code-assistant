package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBin_DummyClientStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("failed to start stream: %v", err)
	}

	// Prepare 10 DummyMessage requests
	var lastDummy *grpcbin.DummyMessage
	for i := 0; i < 10; i++ {
		req := &grpcbin.DummyMessage{
			FString:  "foo",
			FStrings: []string{"bar", "baz"},
			FInt32:   int32(i),
			FInt32S:  []int32{1, 2, 3},
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
			FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
			FBool:    i%2 == 0,
			FBools:   []bool{true, false, true},
			FInt64:   int64(i * 100),
			FInt64S:  []int64{1000, 2000, 3000},
			FBytes:   []byte("bytesdata"),
			FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
			FFloat:   3.14,
			FFloats:  []float32{1.1, 2.2, 3.3},
		}
		lastDummy = req
		if err := stream.Send(req); err != nil {
			t.Fatalf("failed to send DummyMessage #%d: %v", i+1, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("failed to close/send/receive: %v", err)
	}

	if reply == nil {
		t.Error("got nil response from DummyClientStream")
	}

	// Positive response validation (compare the last sent DummyMessage)
	if reply.FString != lastDummy.FString {
		t.Errorf("FString mismatch: got %v, want %v", reply.FString, lastDummy.FString)
	}
	if reply.FInt32 != lastDummy.FInt32 {
		t.Errorf("FInt32 mismatch: got %v, want %v", reply.FInt32, lastDummy.FInt32)
	}
	if reply.FEnum != lastDummy.FEnum {
		t.Errorf("FEnum mismatch: got %v, want %v", reply.FEnum, lastDummy.FEnum)
	}
	if reply.FBool != lastDummy.FBool {
		t.Errorf("FBool mismatch: got %v, want %v", reply.FBool, lastDummy.FBool)
	}
	if reply.FInt64 != lastDummy.FInt64 {
		t.Errorf("FInt64 mismatch: got %v, want %v", reply.FInt64, lastDummy.FInt64)
	}
	if string(reply.FBytes) != string(lastDummy.FBytes) {
		t.Errorf("FBytes mismatch: got %v, want %v", reply.FBytes, lastDummy.FBytes)
	}
	if reply.FFloat != lastDummy.FFloat {
		t.Errorf("FFloat mismatch: got %v, want %v", reply.FFloat, lastDummy.FFloat)
	}

	// (Optionally compare repeated and message fields as well)
}
