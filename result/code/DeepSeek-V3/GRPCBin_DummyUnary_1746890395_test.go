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

func TestDummyUnary(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	testCases := []struct {
		name    string
		request *grpcbin.DummyMessage
	}{
		{
			name: "SimpleRequest",
			request: &grpcbin.DummyMessage{
				FString: "test",
			},
		},
		{
			name: "FullRequest",
			request: &grpcbin.DummyMessage{
				FString:  "test",
				FStrings: []string{"a", "b", "c"},
				FInt32:   42,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:   true,
				FBools:  []bool{true, false, true},
				FInt64:  1234567890,
				FInt64S: []int64{123, 456, 789},
				FBytes:  []byte("bytes"),
				FBytess: [][]byte{[]byte("a"), []byte("b")},
				FFloat:  3.14,
				FFloats: []float32{1.1, 2.2, 3.3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.DummyUnary(ctx, tc.request)
			if err != nil {
				t.Errorf("DummyUnary failed: %v", err)
				return
			}

			if !proto.Equal(resp, tc.request) {
				t.Errorf("Response does not match request.\nGot: %v\nWant: %v", resp, tc.request)
			}
		})
	}
}
