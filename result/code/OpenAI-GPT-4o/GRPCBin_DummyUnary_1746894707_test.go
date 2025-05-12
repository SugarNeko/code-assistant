package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	// Connect to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Construct the request
	req := &grpcbin.DummyMessage{
		FString:   "test_string",
		FStrings:  []string{"string1", "string2"},
		FInt32:    42,
		FInt32S:   []int32{1, 2, 3},
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FEnums:    []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_2},
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		FSubs:     []*grpcbin.DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:     true,
		FBools:    []bool{true, false},
		FInt64:    64,
		FInt64S:   []int64{64, 128},
		FBytes:    []byte("bytes"),
		FBytess:   [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:    3.14,
		FFloats:   []float32{1.1, 2.2},
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the request
	res, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Validate the response
	if res == nil {
		t.Fatal("Response message is nil")
	}

	if res.FString != req.FString {
		t.Errorf("Expected FString %v, got %v", req.FString, res.FString)
	}

	if res.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32 %v, got %v", req.FInt32, res.FInt32)
	}

	log.Printf("Successfully validated response: %v", res)
}
