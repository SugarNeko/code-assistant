package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	req := &grpcbin.DummyMessage{
		FString:   "test",
		FInt32:    123,
		FEnum:     grpcbin.DummyMessage_ENUM_1,
		FSub:      &grpcbin.DummyMessage_Sub{FString: "sub"},
		FBool:     true,
		FInt64:    456,
		FBytes:    []byte("bytes"),
		FFloat:    1.23,
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	res, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	if res.FString != req.FString || res.FInt32 != req.FInt32 || res.FEnum != req.FEnum {
		t.Errorf("Unexpected response: got %v, want %v", res, req)
	}
}
