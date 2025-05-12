package grpcbin_test

import (
	"context"
	"testing"
	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestGRPCBin_DummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "Test String",
		FStrings: []string{"string1", "string2"},
		FInt32: 123,
		FInt32S: []int32{1, 2, 3},
		FEnum: grpcbin.DummyMessage_ENUM_1,
		FEnums: []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub: &grpcbin.DummyMessage_Sub{FString: "Sub String"},
		FSubs: []*grpcbin.DummyMessage_Sub{{FString: "Sub1"}, {FString: "Sub2"}},
		FBool: true,
		FBools: []bool{true, false, true},
		FInt64: 123456789,
		FInt64S: []int64{111, 222, 333},
		FBytes: []byte("Bytes"),
		FBytess: [][]byte{[]byte("Bytes1"), []byte("Bytes2")},
		FFloat: 1.23,
		FFloats: []float32{1.1, 2.2, 3.3},
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	if res.FString != req.FString {
		t.Errorf("Expected response FString %v, got %v", req.FString, res.FString)
	}
	// Add further response validations as needed
}
