package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub"},
		FInt64:   9876543210,
		FBytes:   []byte("hello"),
		FFloat:   1.23,
		FStrings: []string{"a", "b", "c"},
		FBools:   []bool{true, false},
		FInt32s:  []int32{1, 2, 3},
		FInt64s:  []int64{3, 2, 1},
		FBytess:  [][]byte{[]byte("hello"), []byte("world")},
		FFloats:  []float32{1.0, 2.0, 3.0},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		count++
		// Validate response, example checks
		if resp == nil || resp.FString != req.FString {
			t.Errorf("Unexpected response: %v", resp)
		}
	}

	if count != 10 {
		t.Errorf("Expected 10 responses, got %d", count)
	}
}
