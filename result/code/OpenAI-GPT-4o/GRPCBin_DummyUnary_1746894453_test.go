package grpcbin_test

import (
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	request := &grpcbin.DummyMessage{
		FString:   "test",
		FStrings:  []string{"one", "two"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    64,
		FInt64S:   []int64{64, 128},
		FBytes:    []byte{0x01, 0x02},
		FBytess:   [][]byte{{0x01}, {0x02}},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	response, err := client.DummyUnary(context.Background(), request)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	// Validate response
	if response.FString != request.FString {
		t.Errorf("Expected FString %v, got %v", request.FString, response.FString)
	}
	if len(response.FStrings) != len(request.FStrings) {
		t.Errorf("Expected FStrings length %v, got %v", len(request.FStrings), len(response.FStrings))
	}
	// Add further response validation checks as needed
}
