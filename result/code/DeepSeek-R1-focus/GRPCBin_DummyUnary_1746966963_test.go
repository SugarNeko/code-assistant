package grpcbin_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyUnary(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "grpcb.in:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	tests := []struct {
		name    string
		req     *grpcbin.DummyMessage
		validate func(t *testing.T, req *grpcbin.DummyMessage, resp *grpcbin.DummyMessage)
		wantErr bool
	}{
		{
			name: "positive_test_full_request",
			req: &grpcbin.DummyMessage{
				FString:    "test",
				FStrings:   []string{"a", "b"},
				FInt32:     42,
				FInt32S:    []int32{1, 2},
				FEnum:      grpcbin.DummyMessage_ENUM_1,
				FEnums:     []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub:       &grpcbin.DummyMessage_Sub{FString: "sub"},
				FSubs:      []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:      true,
				FBools:     []bool{true, false},
				FInt64:     1234567890,
				FInt64S:    []int64{123, 456},
				FBytes:     []byte("bytes"),
				FBytess:    [][]byte{[]byte("b1"), []byte("b2")},
				FFloat:     3.14,
				FFloats:    []float32{1.1, 2.2},
			},
			validate: func(t *testing.T, req *grpcbin.DummyMessage, resp *grpcbin.DummyMessage) {
				if !reflect.DeepEqual(req, resp) {
					t.Errorf("Response mismatch\nWant: %+v\nGot:  %+v", req, resp)
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.DummyUnary(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Fatalf("DummyUnary() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				tt.validate(t, tt.req, resp)
			}
		})
	}
}
