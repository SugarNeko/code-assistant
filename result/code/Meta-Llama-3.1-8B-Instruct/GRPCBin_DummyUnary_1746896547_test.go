package grpcbin

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"

	"typexampproto/code-assistant/proto/grpcbin"
)

func TestGRPCServer(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	tests := []struct {
		name        string
		request     grpcbin.DummyMessage
		wantResponse grpcbin.DummyMessage
		wantErr     bool
	}{
		{
			name: "ValidUnaryRequest",
			request: grpcbin.DummyMessage{
				FString: "Hello",
				FStrings: []string{"world"},
				FInt32:  123,
				FInt32s:  []int32{1, 2},
				FElem:   grpcbin.DummyMessage_EEnum_1,
				FElems:  []grpcbin.DummyMessage_EEnum{grpcbin.DummyMessage_EEnum_1},
				FSub:    &grpcbin.DummyMessage_Sub{FString: "sub"},
				FSubs:   []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:   true,
				FBools:  []bool{true, false},
				FInt64:  int64(456),
				FInt64s: []int64{3, 4},
				FBytes:  []byte{1, 2, 3},
				FBytess: [][]byte{{4, 5, 6}, {7, 8, 9}},
				FFloat:  3.14,
				FFloats: []float32{1.1, 2.2},
			},
			wantResponse: grpcbin.DummyMessage{
				FString: "Hello",
				FStrings: []string{"world"},
				FInt32:  123,
				FInt32s:  []int32{1, 2},
				FElem:   grpcbin.DummyMessage_EEnum_1,
				FElems:  []grpcbin.DummyMessage_EEnum{grpcbin.DummyMessage_EEnum_1},
				FSub:    &grpcbin.DummyMessage_Sub{FString: "sub"},
				FSubs:   []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
				FBool:   true,
				FBools:  []bool{true, false},
				FInt64:  int64(456),
				FInt64s: []int64{3, 4},
				FBytes:  []byte{1, 2, 3},
				FBytess: [][]byte{{4, 5, 6}, {7, 8, 9}},
				FFloat:  3.14,
				FFloats: []float32{1.1, 2.2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			resp, err := client.DummyUnary(ctx, &tc.request)
			if (err != nil) != tc.wantErr {
				t.Errorf("DummyUnary returned error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(resp, tc.wantResponse) {
				t.Errorf("DummyUnary() = %v, want %v", resp, tc.wantResponse)
			}
		})
	}
}
