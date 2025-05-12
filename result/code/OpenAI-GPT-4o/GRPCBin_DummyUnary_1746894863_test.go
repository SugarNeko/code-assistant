package grpcbin_test

import (
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyUnary(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:   "test",
		FInt32:    123,
		FEnum:     pb.DummyMessage_ENUM_1,
		FSub:      &pb.DummyMessage_Sub{FString: "sub"},
		FBool:     true,
		FInt64:    456,
		FBytes:    []byte("bytes"),
		FFloat:    1.23,
	}

	res, err := client.DummyUnary(context.Background(), req)
	if err != nil {
		t.Fatalf("Error calling DummyUnary: %v", err)
	}

	// Validate client response
	if res.FString != req.FString || res.FBool != req.FBool || res.FInt32 != req.FInt32 {
		t.Fatalf("Client response validation failed. Expected: %v, Got: %v", req, res)
	}

	// Server response validation
	// Add any additional server-side logic checks here, if needed
	log.Printf("Received response: %v", res)
}
