package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FBool:    true,
		FInt64:   1234567890,
		FBytes:   []byte("byte_data"),
		FFloat:   1.23,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("DummyServerStream failed: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		if resp.FString != "test" || resp.FInt32 != 123 {
			t.Errorf("Unexpected response: %v", resp)
		}

		count++
	}

	if count != 10 {
		t.Errorf("Expected 10 responses, got %d", count)
	}
}
