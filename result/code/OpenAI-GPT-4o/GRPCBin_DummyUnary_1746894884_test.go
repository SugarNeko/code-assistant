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
	// Set up a connection to the server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGRPCBinClient(conn)

	// Construct typical request
	req := &pb.DummyMessage{
		FString: "test",
		FInt32:  123,
		FEnum:   pb.DummyMessage_ENUM_1,
		FSub:    &pb.DummyMessage_Sub{FString: "sub_test"},
		FBool:   true,
		FInt64:  456,
		FFloat:  1.23,
	}

	// Set a context with timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the DummyUnary method
	resp, err := client.DummyUnary(ctx, req)
	if err != nil {
		t.Fatalf("DummyUnary failed: %v", err)
	}

	// Validate the server response matches the request
	if resp.FString != req.FString {
		t.Errorf("expected FString %v, got %v", req.FString, resp.FString)
	}
	if resp.FInt32 != req.FInt32 {
		t.Errorf("expected FInt32 %v, got %v", req.FInt32, resp.FInt32)
	}
	if resp.FEnum != req.FEnum {
		t.Errorf("expected FEnum %v, got %v", req.FEnum, resp.FEnum)
	}
	if resp.FSub.FString != req.FSub.FString {
		t.Errorf("expected FSub FString %v, got %v", req.FSub.FString, resp.FSub.FString)
	}
	if resp.FBool != req.FBool {
		t.Errorf("expected FBool %v, got %v", req.FBool, resp.FBool)
	}
	if resp.FInt64 != req.FInt64 {
		t.Errorf("expected FInt64 %v, got %v", req.FInt64, resp.FInt64)
	}
	if resp.FFloat != req.FFloat {
		t.Errorf("expected FFloat %v, got %v", req.FFloat, resp.FFloat)
	}
}
