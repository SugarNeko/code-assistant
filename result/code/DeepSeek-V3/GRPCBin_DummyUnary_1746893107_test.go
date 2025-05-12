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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	testCases := []struct {
		name    string
		request *grpcbin.DummyMessage
		want    *grpcbin.DummyMessage
	}{
		{
			name: "SimpleMessage",
			request: &grpcbin.DummyMessage{
				FString: "test",
				FInt32:  42,
				FBool:   true,
			},
			want: &grpcbin.DummyMessage{
				FString: "test",
				FInt32:  42,
				FBool:   true,
			},
		},
		{
			name: "WithRepeatedFields",
			request: &grpcbin.DummyMessage{
				FStrings: []string{"a", "b", "c"},
				FInt32s:  []int32{1, 2, 3},
				FBools:   []bool{true, false, true},
			},
			want: &grpcbin.DummyMessage{
				FStrings: []string{"a", "b", "c"},
				FInt32s:  []int32{1, 2, 3},
				FBools:   []bool{true, false, true},
			},
		},
		{
			name: "WithSubMessage",
			request: &grpcbin.DummyMessage{
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "submessage",
				},
			},
			want: &grpcbin.DummyMessage{
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "submessage",
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

			if !proto.Equal(resp, tc.want) {
				t.Errorf("Response doesn't match expected:\nGot: %v\nWant: %v", resp, tc.want)
			}
		})
	}
}
