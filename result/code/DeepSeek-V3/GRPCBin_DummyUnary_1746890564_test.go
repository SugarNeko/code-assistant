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

	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
			name: "Basic fields",
			request: &grpcbin.DummyMessage{
				FString:  "test",
				FInt32:   42,
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FBool:    true,
				FInt64:   1234567890,
				FBytes:   []byte("test bytes"),
				FFloat:   3.14,
				FStrings: []string{"a", "b", "c"},
				FInt32s:  []int32{1, 2, 3},
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FBools:   []bool{true, false, true},
				FInt64s:  []int64{123, 456, 789},
				FBytess:  [][]byte{[]byte("a"), []byte("b")},
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
			name:    "Empty message",
			request: &grpcbin.DummyMessage{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.DummyUnary(ctx, tc.request)
			if err != nil {
				t.Fatalf("DummyUnary failed: %v", err)
			}

			if !proto.Equal(tc.request, resp) {
				t.Errorf("Expected response to match request\ngot: %v\nwant: %v", resp, tc.request)
			}
		})
	}
}

func TestServerResponseValidation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "validation test",
	}

	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("Expected FString to be %q, got %q", req.FString, resp.FString)
	}
}
