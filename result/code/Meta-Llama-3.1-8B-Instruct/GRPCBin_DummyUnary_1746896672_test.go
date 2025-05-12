package grpcbin_test

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
็บgrpcConn, err := grpc.Dial("grpcb.in:9000", grpc 銕)
	if err != nil {
		t.Fatal(err)
	}
	defer grpcConn.Close()

	client := pb.NewGRPCBinClient(grpcConn)

	tests := []struct {
		name    string
		msg     pb.DummyMessage
		want    pb.DummyMessage
		wantErr bool
	}{
		{
			name: "valid request",
			msg: pb.DummyMessage{
				FString:    "foo",
				FStrings:   []string{"bar", "baz"},
				FInt32:     42,
				FInt32s:    []int32{1, 2, 3},
				FElem:      pb.DummyMessage_ENUM_1,
				FEnums:     []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1},
				FSub:       pb.DummyMessage_Sub{FString: "sub"},
				FSubs:      []pb.DummyMessage_Sub{{FString: "sub2"}},
				FFBool:     true,
				FFBools:    []bool{true, false},
				FInt64:     64,
				FInt64s:    []int64{1, 2, 3},
				FFBytes:    []byte("hello"),
				FFBytess:   [][]byte{[]byte("hello2")},
				FFlt:       17.5,
				FFFloats:   []float32{1.1, 2.2, 3.3},
			},
			want: pb.DummyMessage{
				FString:    "foo",
				FStrings:   []string{"bar", "baz"},
				FInt32:     42,
				FInt32s:    []int32{1, 2, 3},
				FElem:      pb.DummyMessage_ENUM_1,
				FEnums:     []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1},
				FSub:       pb.DummyMessage_Sub{FString: "sub"},
				FSubs:      []pb.DummyMessage_Sub{{FString: "sub2"}},
				FFBool:     true,
				FFBools:    []bool{true, false},
				FInt64:     64,
				FInt64s:    []int64{1, 2, 3},
				FFBytes:    []byte("hello"),
				FFBytess:   [][]byte{[]byte("hello2")},
				FFFlt:      17.5,
				FFFloats:   []float32{1.1, 2.2, 3.3},
			},
		},
		{
			name: "invalid request - empty strings",
			msg: pb.DummyMessage{
				FString:    "",
				FStrings:   []string{},
				FInt32:     42,
				FInt32s:    []int32{1, 2, 3},
				FElem:      pb.DummyMessage_ENUM_1,
				FEnums:     []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_1},
				FSub:       pb.DummyMessage_Sub{FString: "sub"},
				FSubs:      []pb.DummyMessage_Sub{{FString: "sub2"}},
				FFBool:     true,
				FFBools:    []bool{true, false},
				FInt64:     64,
				FInt64s:    []int64{1, 2, 3},
				FFBytes:    []byte("hello"),
				FFBytess:   [][]byte{[]byte("hello2")},
				FFlt:       17.5,
				FFFloats:   []float32{1.1, 2.2, 3.3},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := client.DummyUnary(context.Background(), &tt.msg)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Call returned error: %v, want nil", err)
				}
				return
			}
			tt.want.FInt64 = 0
			if !reflect.DeepEqual(r, &tt.want) {
				t.Errorf("response = %+v, want %+v", r, tt.want)
			}
		})
	}
}
