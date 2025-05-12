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

	tests := []struct {
		name    string
		request *grpcbin.DummyMessage
		want    *grpcbin.DummyMessage
	}{
		{
			name: "basic request",
			request: &grpcbin.DummyMessage{
				FString:  "test",
				FInt32:   42,
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FBool:    true,
				FInt64:   123456789,
				FBytes:   []byte("test bytes"),
				FFloat:   3.14,
				FStrings: []string{"a", "b", "c"},
				FInt32s:  []int32{1, 2, 3},
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FBools:   []bool{true, false, true},
				FInt64s:  []int64{987654321, 123456789},
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
			want: &grpcbin.DummyMessage{
				FString:  "test",
				FInt32:   42,
				FEnum:    grpcbin.DummyMessage_ENUM_1,
				FBool:    true,
				FInt64:   123456789,
				FBytes:   []byte("test bytes"),
				FFloat:   3.14,
				FStrings: []string{"a", "b", "c"},
				FInt32s:  []int32{1, 2, 3},
				FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
				FBools:   []bool{true, false, true},
				FInt64s:  []int64{987654321, 123456789},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.DummyUnary(ctx, tt.request)
			if err != nil {
				t.Errorf("DummyUnary() error = %v", err)
				return
			}

			if !proto.Equal(got, tt.want) {
				t.Errorf("DummyUnary() got = %v, want %v", got, tt.want)
			}
		})
	}
}
