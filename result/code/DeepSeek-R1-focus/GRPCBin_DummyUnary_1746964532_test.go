package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcbin "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"a", "b"},
		FInt32:   123,
		FInt32S:  []int32{1, 2},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   456,
		FInt64S:  []int64{3, 4},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   1.23,
		FFloats:  []float32{5.6, 7.8},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary RPC failed: %v", err)
	}

	if !reflect.DeepEqual(req, res) {
		t.Errorf("response mismatch\nexpected: %+v\nreceived: %+v", req, res)
	}
}
