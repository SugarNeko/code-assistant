package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString: "test",
		FInt32:  10,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "sub_test"},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	expectedResponseCount := 10
	for i := 0; i < expectedResponseCount; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}

		if resp.FString != req.FString*10 {
			t.Errorf("Expected FString to be %s, got %s", req.FString*10, resp.FString)
		}

		if resp.FInt32 != req.FInt32*10 {
			t.Errorf("Expected FInt32 to be %d, got %d", req.FInt32*10, resp.FInt32)
		}

		if resp.FEnum != req.FEnum {
			t.Errorf("Expected FEnum to be %v, got %v", req.FEnum, resp.FEnum)
		}
	}
}
