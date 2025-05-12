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
			name: "Basic request",
			request: &grpcbin.DummyMessage{
				FString:  "test",
				FStrings: []string{"a", "b", "c"},
				FInt32:   123,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub string",
				},
				FBool:   true,
				FBools:  []bool{true, false, true},
				FInt64:  456,
				FInt64S: []int64{4, 5, 6},
				FBytes:  []byte("test bytes"),
			},
		},
		{
			name: "Empty request",
			request: &grpcbin.DummyMessage{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := c.DummyUnary(ctx, tc.request)
			if err != nil {
				t.Errorf("DummyUnary failed: %v", err)
				return
			}

			if !proto.Equal(tc.request, resp) {
				t.Errorf("Expected response to match request, got %v, want %v", resp, tc.request)
			}
		})
	}
}

func TestServerResponseValidation(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := grpcbin.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &grpcbin.DummyMessage{
		FString: "validation test",
		FInt32:  999,
	}

	resp, err := c.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if resp.FString != req.FString {
		t.Errorf("FString mismatch, got %q, want %q", resp.FString, req.FString)
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch, got %d, want %d", resp.FInt32, req.FInt32)
	}
}
