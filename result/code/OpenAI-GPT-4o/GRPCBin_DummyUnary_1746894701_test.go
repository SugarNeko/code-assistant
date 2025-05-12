package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	request := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub_test"},
		FBool:    true,
		FInt64:   321,
		FBytes:   []byte("byte_test"),
		FFloat:   1.23,
	}

	resp, err := client.DummyUnary(context.Background(), request)
	if err != nil {
		t.Fatalf("DummyUnary call failed: %v", err)
	}

	if resp.FString != request.FString || resp.FInt32 != request.FInt32 || resp.FEnum != request.FEnum {
		t.Errorf("Response does not match request: got %v, want %v", resp, request)
	}
}
