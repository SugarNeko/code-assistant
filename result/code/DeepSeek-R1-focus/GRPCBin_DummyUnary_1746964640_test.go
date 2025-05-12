package grpcbin_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary_Positive(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"a", "b"},
		FInt32:   123,
		FInt32s:  []int32{1, 2, 3},
		FEnum:    grpcbin.Enum_ENUM_1,
		FEnums:   []grpcbin.Enum{grpcbin.Enum_ENUM_0, grpcbin.Enum_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:    []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   456,
		FInt64s:  []int64{4, 5, 6},
		FBytes:   []byte("bytes"),
		FBytess:  [][]byte{[]byte("b1"), []byte("b2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if !reflect.DeepEqual(resp, req) {
		t.Error("Response does not match request")
	}
}
