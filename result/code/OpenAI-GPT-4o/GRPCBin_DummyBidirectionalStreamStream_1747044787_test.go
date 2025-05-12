package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	dummyMessage := &pb.DummyMessage{
		FString:  "hello",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub"},
		FBool:    true,
		FInt64:   456,
		FBytes:   []byte("bytes"),
		FFloat:   3.14,
	}

	if err := stream.Send(dummyMessage); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	if resp.FString != dummyMessage.FString ||
		resp.FInt32 != dummyMessage.FInt32 ||
		resp.FEnum != dummyMessage.FEnum ||
		resp.FSub.FString != dummyMessage.FSub.FString ||
		resp.FBool != dummyMessage.FBool ||
		resp.FInt64 != dummyMessage.FInt64 ||
		resp.FBytes[0] != dummyMessage.FBytes[0] ||
		resp.FFloat != dummyMessage.FFloat {
		t.Fatal("Response did not match request data")
	}
}
