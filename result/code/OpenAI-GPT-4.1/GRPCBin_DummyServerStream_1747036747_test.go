package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Construct a typical request
	req := &pb.DummyMessage{
		FString:   "hello",
		FStrings:  []string{"a", "b", "c"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     pb.DummyMessage_ENUM_1,
		FEnums:    []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FSub:      &pb.DummyMessage_Sub{FString: "submessage"},
		FSubs:     []*pb.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    1001,
		FInt64S:   []int64{10, 11, 12},
		FBytes:    []byte{0x0a, 0x0b},
		FBytess:   [][]byte{{0x01, 0x02}, {0x03, 0x04}},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream RPC failed: %v", err)
	}

	// We expect 10 responses, each equal to the request
	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("failed to receive message %d: %v", i, err)
		}

		if resp.FString != req.FString {
			t.Errorf("response FString mismatch at msg %d: got %q, want %q", i, resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("response FStrings length mismatch at msg %d: got %v, want %v", i, resp.FStrings, req.FStrings)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response FInt32 mismatch at msg %d: got %v, want %v", i, resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("response FEnum mismatch at msg %d: got %v, want %v", i, resp.FEnum, req.FEnum)
		}
		if resp.FSub == nil || resp.FSub.FString != req.FSub.FString {
			t.Errorf("response FSub mismatch at msg %d: got %v, want %v", i, resp.FSub, req.FSub)
		}
		if resp.FBool != req.FBool {
			t.Errorf("response FBool mismatch at msg %d: got %v, want %v", i, resp.FBool, req.FBool)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("response FInt64 mismatch at msg %d: got %v, want %v", i, resp.FInt64, req.FInt64)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("response FBytes mismatch at msg %d: got %v, want %v", i, resp.FBytes, req.FBytes)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("response FFloat mismatch at msg %d: got %v, want %v", i, resp.FFloat, req.FFloat)
		}
	}

	_, err = stream.Recv()
	if err == nil {
		t.Fatalf("expected end of stream, but got more messages")
	}
}
