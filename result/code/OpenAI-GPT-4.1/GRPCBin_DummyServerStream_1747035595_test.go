package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyServerStream_Positive(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "foo",
		FStrings:  []string{"foo", "bar"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "subfoo"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false, true},
		FInt64:    424242,
		FInt64S:   []int64{101, 202},
		FBytes:    []byte("hello"),
		FBytess:   [][]byte{[]byte("a"), []byte("b")},
		FFloat:    3.14,
		FFloats:   []float32{1.23, 4.56},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	expectedCount := 10
	for i := 0; i < expectedCount; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("error receiving from stream at message %d: %v", i, err)
		}
		// Basic validation: check some fields are echoed back as expected
		if resp.FString != req.FString {
			t.Errorf("response[%d]: got FString %q, want %q", i, resp.FString, req.FString)
		}
		if len(resp.FStrings) != len(req.FStrings) {
			t.Errorf("response[%d]: got FStrings len %d, want %d", i, len(resp.FStrings), len(req.FStrings))
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("response[%d]: got FInt32 %d, want %d", i, resp.FInt32, req.FInt32)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("response[%d]: got FEnum %v, want %v", i, resp.FEnum, req.FEnum)
		}
		if resp.FBool != req.FBool {
			t.Errorf("response[%d]: got FBool %v, want %v", i, resp.FBool, req.FBool)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("response[%d]: got FFloat %v, want %v", i, resp.FFloat, req.FFloat)
		}
	}

	_, err = stream.Recv()
	if err == nil {
		t.Errorf("expected EOF after %d stream responses", expectedCount)
	}
}
