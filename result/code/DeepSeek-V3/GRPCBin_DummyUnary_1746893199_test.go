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
			name: "basic fields",
			request: &grpcbin.DummyMessage{
				FString:  "test",
				FStrings: []string{"a", "b"},
				FInt32:   42,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FBool:    true,
				FBools:   []bool{true, false, true},
				FInt64:   123456789,
				FInt64S:  []int64{987654321, 123456789},
				FBytes:   []byte("test bytes"),
				FBytess:  [][]byte{[]byte("a"), []byte("b")},
				FFloat:   3.14,
				FFloats:  []float32{1.1, 2.2, 3.3},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub string",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
			},
			want: &grpcbin.DummyMessage{
				FString:  "test",
				FStrings: []string{"a", "b"},
				FInt32:   42,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FBool:    true,
				FBools:   []bool{true, false, true},
				FInt64:   123456789,
				FInt64S:  []int64{987654321, 123456789},
				FBytes:   []byte("test bytes"),
				FBytess:  [][]byte{[]byte("a"), []byte("b")},
				FFloat:   3.14,
				FFloats:  []float32{1.1, 2.2, 3.3},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub string",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
			},
		},
		{
			name:    "empty message",
			request: &grpcbin.DummyMessage{},
			want:    &grpcbin.DummyMessage{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := c.DummyUnary(ctx, tc.request)
			if err != nil {
				t.Errorf("DummyUnary failed: %v", err)
			}

			if !proto.Equal(resp, tc.want) {
				t.Errorf("Response doesn't match expected\nGot: %v\nWant: %v", resp, tc.want)
			}
		})
	}
}
