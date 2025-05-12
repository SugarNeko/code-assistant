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
			name: "Simple request",
			request: &grpcbin.DummyMessage{
				FString: "test string",
			},
		},
		{
			name: "Complex request",
			request: &grpcbin.DummyMessage{
				FString:  "test string",
				FStrings: []string{"one", "two", "three"},
				FInt32:   42,
				FInt32s:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "sub string",
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
			resp, err := client.DummyUnary(ctx, tc.request)
			if err != nil {
				t.Errorf("DummyUnary failed: %v", err)
				return
			}

			if !proto.Equal(resp, tc.request) {
				t.Errorf("Expected response to match request. Got: %v, Want: %v", resp, tc.request)
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
		t.Errorf("FString mismatch. Got: %q, Want: %q", resp.FString, req.FString)
	}

	if len(resp.FStrings) != len(req.FStrings) {
		t.Errorf("FStrings length mismatch. Got: %d, Want: %d", len(resp.FStrings), len(req.FStrings))
	}

	if resp.FInt32 != req.FInt32 {
		t.Errorf("FInt32 mismatch. Got: %d, Want: %d", resp.FInt32, req.FInt32)
	}

	if resp.FEnum != req.FEnum {
		t.Errorf("FEnum mismatch. Got: %v, Want: %v", resp.FEnum, req.FEnum)
	}

	if resp.FSub != nil && req.FSub != nil {
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("FSub.FString mismatch. Got: %q, Want: %q", resp.FSub.FString, req.FSub.FString)
		}
	}
}
