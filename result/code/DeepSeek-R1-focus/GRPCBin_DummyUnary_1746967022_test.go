package grpcbin_test

import (
	"context"
	"reflect"
	"testing"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var client grpcbin.GRPCBinClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client = grpcbin.NewGRPCBinClient(conn)
	m.Run()
}

func TestDummyUnary_Positive(t *testing.T) {
	req := &grpcbin.DummyMessage{
		FString:    "test",
		FStrings:   []string{"a", "b"},
		FInt32:     42,
		FInt32s:    []int32{1, 2},
		FEnum:      grpcbin.DummyMessage_ENUM_1,
		FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSub:       &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:      true,
		FBools:     []bool{true, false},
		FInt64:     123456789,
		FInt64s:    []int64{987654321},
		FBytes:     []byte("bytes"),
		FBytess:    [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:     3.14,
		FFloats:    []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if !reflect.DeepEqual(req, resp) {
		t.Errorf("Response doesn't match request\nWant: %+v\nGot:  %+v", req, resp)
	}
}
