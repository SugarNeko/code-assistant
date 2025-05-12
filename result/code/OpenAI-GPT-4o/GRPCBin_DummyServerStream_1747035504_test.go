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
		FString: "test",
		FInt32:  42,
		FEnum:   pb.DummyMessage_ENUM_1,
		FBool:   true,
		FInt64:  1000,
		FFloat:  1.234,
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to request: %v", err)
	}

	expectedResponseCount := 10
	responseCount := 0

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		// Validate response
		if resp.FString != req.FString {
			t.Errorf("Expected FString %s, got %s", req.FString, resp.FString)
		}

		if resp.FInt32 != req.FInt32 {
			t.Errorf("Expected FInt32 %d, got %d", req.FInt32, resp.FInt32)
		}

		if resp.FEnum != req.FEnum {
			t.Errorf("Expected FEnum %v, got %v", req.FEnum, resp.FEnum)
		}

		if resp.FInt64 != req.FInt64 {
			t.Errorf("Expected FInt64 %d, got %d", req.FInt64, resp.FInt64)
		}

		if resp.FFloat != req.FFloat {
			t.Errorf("Expected FFloat %f, got %f", req.FFloat, resp.FFloat)
		}

		responseCount++
	}

	if responseCount != expectedResponseCount {
		t.Errorf("Expected %d responses, got %d", expectedResponseCount, responseCount)
	}
}
