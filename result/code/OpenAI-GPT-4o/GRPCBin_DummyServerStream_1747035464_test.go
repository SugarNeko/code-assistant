package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	req := &pb.DummyMessage{
		FString:  "Test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "SubTest"},
		FBool:    true,
		FInt64:   1234567890,
		FFloat:   1.23,
		FBools:   []bool{true, false},
		FStrings: []string{"str1", "str2"},
		FInt32s:  []int32{1, 2, 3},
		FEnums:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
		FInt64s:  []int64{123, 456},
		FBytes:   []byte("byteContent"),
		FFloats:  []float32{1.0, 2.0},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}
	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}
		if resp == nil || resp.FString != req.FString {
			t.Errorf("Unexpected response, got: %v, want: %v", resp.FString, req.FString)
		}
	}

	// Verify trailer
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to close stream: %v", err)
	}
	_, errTrailer := stream.Recv()
	if status.Code(errTrailer) != codes.OutOfRange {
		t.Errorf("Expected codes.OutOfRange error, got: %v", errTrailer)
	}
}
