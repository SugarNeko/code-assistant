package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestGRPCBinDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	req := &pb.DummyMessage{
		FString: "test",
		FInt32:  42,
		FBool:   true,
	}

	stream, err := client.DummyServerStream(ctx, req)
	if err != nil {
		t.Fatalf("Error calling DummyServerStream: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		
		if resp.FString != req.FString {
			t.Errorf("Expected FString: %v, got: %v", req.FString, resp.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Expected FInt32: %v, got: %v", req.FInt32, resp.FInt32)
		}
		if resp.FBool != req.FBool {
			t.Errorf("Expected FBool: %v, got: %v", req.FBool, resp.FBool)
		}
		count++
	}
	
	if count != 10 {
		t.Errorf("Expected 10 responses, got %d", count)
	}
}
