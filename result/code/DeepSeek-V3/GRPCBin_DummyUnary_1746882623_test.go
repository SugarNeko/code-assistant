package grpcbin_test

import (
	"context"
	"testing"

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

	tests := []struct {
		name    string
		req     *grpcbin.DummyMessage
		want    *grpcbin.DummyMessage
		wantErr bool
	}{
		{
			name: "basic request",
			req: &grpcbin.DummyMessage{
				FString:  "test",
				FStrings: []string{"a", "b", "c"},
				FInt32:   42,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "substring",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:   true,
				FBools:  []bool{true, false, true},
				FInt64:  1234567890,
				FInt64S: []int64{9876543210, 1234567890},
				FBytes:  []byte("test bytes"),
				FBytess: [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:  3.14,
				FFloats: []float32{1.1, 2.2, 3.3},
			},
			want: &grpcbin.DummyMessage{
				FString:  "test",
				FStrings: []string{"a", "b", "c"},
				FInt32:   42,
				FInt32S:  []int32{1, 2, 3},
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FSub: &grpcbin.DummyMessage_Sub{
					FString: "substring",
				},
				FSubs: []*grpcbin.DummyMessage_Sub{
					{FString: "sub1"},
					{FString: "sub2"},
				},
				FBool:   true,
				FBools:  []bool{true, false, true},
				FInt64:  1234567890,
				FInt64S: []int64{9876543210, 1234567890},
				FBytes:  []byte("test bytes"),
				FBytess: [][]byte{[]byte("bytes1"), []byte("bytes2")},
				FFloat:  3.14,
				FFloats: []float32{1.1, 2.2, 3.3},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.DummyUnary(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyUnary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("DummyUnary() = %v, want %v", got, tt.want)
			}
		})
	}
}
