package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Error creating stream: %v", err)
	}

	// Send 10 DummyMessages
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString:  "test",
			FInt32:   int32(i),
			FEnum:    pb.DummyMessage_ENUM_1,
			FSub:     &pb.DummyMessage_Sub{FString: "sub"},
			FBool:    true,
			FInt64:   int64(i),
			FFloat:   float32(i),
			FBytes:   []byte{1, 2, 3},
			FStrings: []string{"string1", "string2"},
		}

		if err := stream.Send(msg); err != nil {
			t.Fatalf("Error sending message: %v", err)
		}
	}

	// Close the send direction of the stream
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Error closing stream: %v", err)
	}

	// Receive last DummyMessage
	response, err := stream.Recv()
	if err != nil {
		t.Fatalf("Error receiving response: %v", err)
	}

	// Validate Response
	expected := "test"
	if response.FString != expected {
		t.Errorf("Expected response FString to be %v, got %v", expected, response.FString)
	}

	// Additional response validation...
	if !response.FBool {
		t.Errorf("Expected response FBool to be true, got false")
	}
}
