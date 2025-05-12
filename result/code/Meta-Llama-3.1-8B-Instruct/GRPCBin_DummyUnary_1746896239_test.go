package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	t.Parallel()

	Array0 := []string{"a", "b", "c"}
	Array1 := []bool{true, false}
	Array2 := []int64{1, 2, 3}
	Array3 := []float32{1.0, 2.0}
	Array4 := []byte{1, 2, 3}

	tests := []struct {
		in    *pb.DummyMessage
		expect *pb.DummyMessage
	}{
		{
			in: &pb.DummyMessage{
				FString:          "string",
				FStrings:         Array0,
				FInt32:           256,
				FInt32s:          Array1,
				FElem:            pb.DummyMessage_ENUM_1,
				FElems:           [3]pb.DummyMessage_ENUM{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
				FSub:             &pb.DummyMessage_Sub{FString: "string"},
				FSubs:            Array2,
				FBool:            false,
				FBools:           Array3,
				FFloat:           255,
				FFloats:          Array4,
				FInt64:           int64(time.Now().UnixNano()) + 1,
				FInt64s:          Array5,
				FFBytes:          []byte("a"),
				FBytess:          Array6,
			},
			expect: &pb.DummyMessage{
				FString:          "string",
				FStrings:         Array0,
				FInt32:           256,
				FInt32s:          Array1,
				FElem:            pb.DummyMessage_ENUM_1,
				FElems:           [3]pb.DummyMessage_ENUM{pb.DummyMessage_ENUM_1, pb.DummyMessage_ENUM_2, pb.DummyMessage_ENUM_0},
				FSub:             &pb.DummyMessage_Sub{FString: "string"},
				FSubs:            Array2,
				FBool:            false,
				FBools:           Array3,
				FFloat:           255,
				FFloats:          Array4,
				FInt64:           int64(time.Now().UnixNano()) + 1,
				FInt64s:          Array5,
				FFBytes:          []byte("a"),
				FBytess:          Array6,
			},
		},
		// TODO: add more test cases
	}

	clientConn, err := grpc.DialContext(context.Background(), "grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer clientConn.Close()

	client := pb.NewGRPCBinClient(clientConn)

	for _, tt := range tests {
		t.Run(tt.expect, func(t *testing.T) {
			res, err := client.DummyUnary(context.Background(), tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(res, tt.expect) {
				t.Errorf("-want  %+v,\n-got %+v", tt.expect, res)
			}
		})
	}
}

var Array5 = []int64{1, 2, 3, 4, 5}
var Array6 = [][]byte{[]byte("a"), []byte("b"), []byte("c")}
