package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "path/to/proto/grpcbin" // Update with your actual path

	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	ctx := context.Background()

	stream, err := client.DummyBidirectionalStreamStream(ctx)
	if err != nil {
		t.Fatalf("failed to open stream: %v", err)
	}

	// Construct a sample request as specified by the proto definition
	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    pb.DummyMessage_ENUM_1,
		FBool:    true,
		FInt64:   1234567890,
		FFloat:   1.23,
		FStrings: []string{"str1", "str2"},
		FSub:     &pb.DummyMessage_Sub{FString: "sub"},
	}

	// Send the request to the server
	if err := stream.Send(req); err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	// Receive and validate the response
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive response: %v", err)
	}

	if resp.FString != req.FString || resp.FInt32 != req.FInt32 || resp.FBool != req.FBool || resp.FInt64 != req.FInt64 || resp.FFloat != req.FFloat {
		t.Error("the response does not match the request")
	}

	if len(resp.FStrings) != len(req.FStrings) {
		t.Error("the response strings do not match the request strings")
	}

	if resp.FSub.FString != req.FSub.FString {
		t.Error("the response sub string does not match the request sub string")
	}
}
