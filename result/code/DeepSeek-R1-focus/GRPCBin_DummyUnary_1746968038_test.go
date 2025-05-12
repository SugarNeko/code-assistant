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

func TestDummyUnary(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	t.Run("PositiveTest", func(t *testing.T) {
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
			FBytes:    []byte("test bytes"),
			FBytess:   [][]byte{[]byte("b1"), []byte("b2")},
			FFloat:    3.14,
			FFloats:   []float32{1.1, 2.2},
		}

		resp, err := client.DummyUnary(ctx, req)
		if err != nil {
			t.Errorf("DummyUnary failed: %v", err)
		}

		if !reflect.DeepEqual(req.FBytes, resp.FBytes) {
			t.Error("Bytes fields mismatch")
		}

		req.FBytes, resp.FBytes = nil, nil
		if !reflect.DeepEqual(req, resp) {
			t.Error("Response does not match request")
		}
	})
}
