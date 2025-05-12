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
	}{
		{
			name: "SimpleRequest",
			request: &grpcbin.DummyMessage{
				FString: "test",
			},
		},
		{
			name: "ComplexRequest",
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
				FBool:   true,
				FBools:  []bool{true, false, true},
				FInt64:  1234567890,
				FBytes:  []byte("test bytes"),
				FFloat:  3.14,
				FFloats: []float32{1.1, 2.2, 3.3},
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
				t.Errorf("Expected response to match request, got %v, want %v", resp, tc.request)
			}
		})
	}
}
