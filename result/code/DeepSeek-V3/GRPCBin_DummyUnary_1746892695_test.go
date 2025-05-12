package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

const (
	address = "grpcb.in:9000"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	testCases := []struct {
		name    string
		request *grpcbin.DummyMessage
	}{
		{
			name: "Basic request",
			request: &grpcbin.DummyMessage{
				FString:  "test",
				FInt32:   42,
				FBool:    true,
				FInt64:   1234567890,
				FFloat:   3.14,
				FBytes:   []byte("test bytes"),
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FStrings: []string{"one", "two", "three"},
				FInt32s:  []int32{1, 2, 3},
				FBools:   []bool{true, false, true},
				FInt64s:  []int64{123, 456, 789},
				FFloats:  []float32{1.1, 2.2, 3.3},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub message",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := c.DummyUnary(ctx, tc.request)
			if err != nil {
				t.Errorf("DummyUnary failed: %v", err)
				return
			}

			if !proto.Equal(resp, tc.request) {
				t.Errorf("Response doesn't match request. Got %v, want %v", resp, tc.request)
			}
		})
	}
}
