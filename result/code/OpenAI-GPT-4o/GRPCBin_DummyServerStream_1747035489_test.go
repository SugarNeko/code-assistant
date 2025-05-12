package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
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
		FInt32:   42,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "subtest"},
		FBool:    true,
		FInt64:   123456789,
		FBytes:   []byte("byteTest"),
		FFloat:   3.14,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Error calling DummyServerStream: %v", err)
	}

	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Error receiving from stream: %v", err)
		}

		if resp.FString != req.FString {
			t.Errorf("Expected FString %v, got %v", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32*10 {
			t.Errorf("Expected FInt32 %v, got %v", req.FInt32*10, resp.FInt32)
		}
	}
}
