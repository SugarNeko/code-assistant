package grpcbin_test

import (
	"context"
	"testing"
	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"test1", "test2"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   12345678,
		FInt64S:  []int64{10, 20, 30},
		FBytes:   []byte{0x1, 0x2},
		FBytess:  [][]byte{{0x3}, {0x4}},
		FFloat:   1.23,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	if res.FString != req.FString {
		t.Errorf("Expected FString %v, got %v", req.FString, res.FString)
	}
	// Additional validation checks for other fields if needed
}
