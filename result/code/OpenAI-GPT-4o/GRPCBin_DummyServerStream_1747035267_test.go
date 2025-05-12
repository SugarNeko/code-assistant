package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyServerStream(t *testing.T) {
	// Set up a connection to the server with a 15-second timeout.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	req := &pb.DummyMessage{
		FString:  "test",
		FInt32:   1,
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub_test"},
		FBool:    true,
		FInt64:   123,
		FFloat:   1.23,
		FBools:   []bool{true, false},
		FEnuns:   []pb.DummyMessage_Enum{pb.DummyMessage_ENUM_0, pb.DummyMessage_ENUM_2},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	count := 0
	for {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive response: %v", err)
		}
		count++
		if resp.FString != req.FString {
			t.Fatalf("Response FString mismatch: got %v, want %v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32*10 {
			t.Fatalf("Response FInt32 mismatch: got %v, want %v", resp.FInt32, req.FInt32*10)
		}
		if count >= 10 {
			break
		}
	}
}
