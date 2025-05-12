package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FStrings: []string{"a", "b", "c"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{FString: "sub1"},
			{FString: "sub2"},
		},
		FBool:   true,
		FBools:  []bool{true, false, true},
		FInt64:  1234567890,
		FInt64S: []int64{987654321, 123456789},
		FBytes:  []byte("test bytes"),
		FBytess: [][]byte{[]byte("a"), []byte("b")},
		FFloat:  3.14,
		FFloats: []float32{1.1, 2.2, 3.3},
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	assert.Equal(t, req.FString, resp.FString)
	assert.Equal(t, req.FStrings, resp.FStrings)
	assert.Equal(t, req.FInt32, resp.FInt32)
	assert.Equal(t, req.FInt32S, resp.FInt32S)
	assert.Equal(t, req.FEnum, resp.FEnum)
	assert.Equal(t, req.FEnums, resp.FEnums)
	assert.Equal(t, req.FSub.FString, resp.FSub.FString)
	assert.Equal(t, len(req.FSubs), len(resp.FSubs))
	for i := range req.FSubs {
		assert.Equal(t, req.FSubs[i].FString, resp.FSubs[i].FString)
	}
	assert.Equal(t, req.FBool, resp.FBool)
	assert.Equal(t, req.FBools, resp.FBools)
	assert.Equal(t, req.FInt64, resp.FInt64)
	assert.Equal(t, req.FInt64S, resp.FInt64S)
	assert.Equal(t, req.FBytes, resp.FBytes)
	assert.Equal(t, req.FBytess, resp.FBytess)
	assert.Equal(t, req.FFloat, resp.FFloat)
	assert.Equal(t, req.FFloats, resp.FFloats)
}
