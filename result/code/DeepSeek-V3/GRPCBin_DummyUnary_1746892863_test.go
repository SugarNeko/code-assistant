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
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	testCases := []struct {
		name    string
		request *grpcbin.DummyMessage
	}{
		{
			name: "basic request",
			request: &grpcbin.DummyMessage{
				FString:  "test",
				FStrings: []string{"a", "b", "c"},
				FInt32:   42,
				FInt32s:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub string",
				},
				FBool:   true,
				FBools:  []bool{true, false, true},
				FInt64:  123456789,
				FBytes:  []byte("test bytes"),
				FFloat:  3.14,
				FFloats: []float32{1.1, 2.2, 3.3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.DummyUnary(ctx, tc.request)
			if err != nil {
				t.Fatalf("DummyUnary failed: %v", err)
			}

			if !proto.Equal(resp, tc.request) {
				t.Errorf("expected response to equal request, got diff: %v", resp)
			}
		})
	}
}
