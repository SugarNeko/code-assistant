package grpcbin_test

import (
	"context"
	"testing"
	"time"
	"reflect"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary_Positive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:   "test",
		FStrings:  []string{"a", "b"},
		FInt32:    42,
		FInt32S:   []int32{1, 2},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    123456789,
		FInt64S:   []int64{987654321, 123456789},
		FBytes:    []byte{0x01, 0x02},
		FBytess:   [][]byte{{0x03}, {0x04}},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if !reflect.DeepEqual(req.FString, resp.FString) ||
		!reflect.DeepEqual(req.FStrings, resp.FStrings) ||
		!reflect.DeepEqual(req.FInt32, resp.FInt32) ||
		!reflect.DeepEqual(req.FInt32S, resp.FInt32S) ||
		!reflect.DeepEqual(req.FEnum, resp.FEnum) ||
		!reflect.DeepEqual(req.FEnums, resp.FEnums) ||
		!reflect.DeepEqual(req.FSub, resp.FSub) ||
		!reflect.DeepEqual(req.FSubs, resp.FSubs) ||
		!reflect.DeepEqual(req.FBool, resp.FBool) ||
		!reflect.DeepEqual(req.FBools, resp.FBools) ||
		!reflect.DeepEqual(req.FInt64, resp.FInt64) ||
		!reflect.DeepEqual(req.FInt64S, resp.FInt64S) ||
		!reflect.DeepEqual(req.FBytes, resp.FBytes) ||
		!reflect.DeepEqual(req.FBytess, resp.FBytess) ||
		!reflect.DeepEqual(req.FFloat, resp.FFloat) ||
		!reflect.DeepEqual(req.FFloats, resp.FFloats) {
		t.Error("Response data does not match request")
	}
}
