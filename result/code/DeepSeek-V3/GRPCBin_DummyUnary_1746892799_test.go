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
		req     *grpcbin.DummyMessage
		want    *grpcbin.DummyMessage
		wantErr bool
	}{
		{
			name: "basic request",
			req: &grpcbin.DummyMessage{
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
				FSubs: []*grpcbin.DummyMessage.Sub{
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
			},
			wantErr: false,
		},
		{
			name: "empty request",
			req:  &grpcbin.DummyMessage{},
			want: &grpcbin.DummyMessage{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.DummyUnary(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DummyUnary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Compare only the fields we expect to be echoed back
				if got.FString != tt.want.FString {
					t.Errorf("DummyUnary() FString = %v, want %v", got.FString, tt.want.FString)
				}
				if got.FInt32 != tt.want.FInt32 {
					t.Errorf("DummyUnary() FInt32 = %v, want %v", got.FInt32, tt.want.FInt32)
				}
				if got.FEnum != tt.want.FEnum {
					t.Errorf("DummyUnary() FEnum = %v, want %v", got.FEnum, tt.want.FEnum)
				}
				if got.FBool != tt.want.FBool {
					t.Errorf("DummyUnary() FBool = %v, want %v", got.FBool, tt.want.FBool)
				}
				if got.FInt64 != tt.want.FInt64 {
					t.Errorf("DummyUnary() FInt64 = %v, want %v", got.FInt64, tt.want.FInt64)
				}
				if string(got.FBytes) != string(tt.want.FBytes) {
					t.Errorf("DummyUnary() FBytes = %v, want %v", got.FBytes, tt.want.FBytes)
				}
				if got.FFloat != tt.want.FFloat {
					t.Errorf("DummyUnary() FFloat = %v, want %v", got.FFloat, tt.want.FFloat)
				}

				// For complex fields, use proto.Equal for deep comparison
				if !proto.Equal(got.FSub, tt.want.FSub) {
					t.Errorf("DummyUnary() FSub = %v, want %v", got.FSub, tt.want.FSub)
				}
			}
		})
	}
}
