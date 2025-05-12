package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestGRPCBinService(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	sendErr := stream.Send(&grpcbin.DummyMessage{
		FString: "test",
		FInt32:  42,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "subtest"},
		FBool:   true,
		FInt64:  64,
	})
	if sendErr != nil {
		t.Fatalf("Failed to send message: %v", sendErr)
	}

	res, recvErr := stream.Recv()
	if recvErr != nil {
		t.Fatalf("Failed to receive message: %v", recvErr)
	}

	if res.FString != "test" || res.FInt32 != 42 || res.FEnum != grpcbin.DummyMessage_ENUM_1 || res.FSub.FString != "subtest" || res.FBool != true || res.FInt64 != 64 {
		t.Fatalf("Unexpected response: %v", res)
	}

	log.Printf("Received message: %v", res)
}
