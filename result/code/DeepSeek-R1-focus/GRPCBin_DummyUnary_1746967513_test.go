package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	ctx := context.Background()

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"a", "b"},
		FInt32:   123,
		FInt32s:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   456,
		FInt64s:  []int64{4, 5, 6},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   1.5,
		FFloats:  []float32{1.0, 2.0},
	}

	res, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary RPC failed: %v", err)
	}

	if !reflect.DeepEqual(req, res) {
		t.Errorf("Response does not match request\nExpected: %+v\nReceived: %+v", req, res)
	}
}
