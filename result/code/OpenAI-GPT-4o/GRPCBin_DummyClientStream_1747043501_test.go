package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "path/to/your/project/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Could not open stream: %v", err)
	}

	expectedMessage := &pb.DummyMessage{
		FString: "test",
		FInt32:  1,
	}

	for i := 0; i < 10; i++ {
		if err := stream.Send(expectedMessage); err != nil {
			t.Fatalf("Failed to send message %d: %v", i, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Error during stream close: %v", err)
	}

	if reply.FString != expectedMessage.FString {
		t.Errorf("Expected FString to be %s, but got %s", expectedMessage.FString, reply.FString)
	}
	if reply.FInt32 != expectedMessage.FInt32 {
		t.Errorf("Expected FInt32 to be %d, but got %d", expectedMessage.FInt32, reply.FInt32)
	}
}
