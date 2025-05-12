package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	// Connect to the gRPC server
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Construct the request
	req := &pb.DummyMessage{
		FString:  "test",
		FStrings: []string{"test1", "test2"},
		FInt32:   123,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    pb.DummyMessage_ENUM_1,
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_2},
		FSub: &pb.DummyMessage_Sub{
			FString: "subTest",
		},
		FSubs: []*pb.DummyMessage_Sub{
			{FString: "subTest1"},
		},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   123456789,
		FInt64S:  []int64{987654321, 123123123},
		FBytes:   []byte("testBytes"),
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:   0.12345,
		FFloats:  []float32{1.234, 5.678},
	}

	// Test the response
	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Errorf("DummyUnary failed: %v", err)
	}

	// Validate the response
	if resp.FString != req.FString {
		t.Errorf("Expected FString %v, got %v", req.FString, resp.FString)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("Expected FInt32 %v, got %v", req.FInt32, resp.FInt32)
	}
	// Add more detailed checks for each field as required
}
