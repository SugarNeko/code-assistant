package grpcbin

import (
	"context"
	"testing"
	"log"

	"google.golang.org/grpc"
)

const (
	grpcServerAddr = "grpcb.in:9000"
)

func TestDummyUnary(t *testing.T) {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := NewGRPCBinClient(conn)

	// Construct a dummy request
	req := &DummyMessage{
		FString: "Hello",
		FStrings: []string{"World", "again"},
		FInt32: 42,
		FInt32s: []int32{1, 2, 3},
		FEnum: Enum.Enum_1,
		FEnums: []Enum{Enum.Enum_0, Enum.Enum_2},
		FSub: &DummyMessage_Sub{
			FString: "Sub",
		},
		FSubs: []*DummyMessage_Sub{
			{FString: "Sub1"},
			{FString: "Sub2"},
		},
		FBool: true,
		FBools: []bool{false, true},
		FInt64: 123,
		FInt64s: []int64{456, 789},
		FBytes: []byte("Bytes"),
		FBytess: [][]byte{[]byte("Bytess1"), []byte("Bytess2")},
		FFloat: 1.23,
		FFloats: []float32{4.56, 7.89},
	}

	// Perform the unary RPC call
	resp, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Errorf("DummyUnary RPC call failed: %v", err)
	}

	// Validate the response
	if resp == nil {
		t.Errorf("response is nil")
	}
	if resp.FString != req.FString {
		t.Errorf("FString mismatch: %s != %s", resp.FString, req.FString)
	}
	// ... validate other fields as well
}
