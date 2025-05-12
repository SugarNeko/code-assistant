package grpcbintest

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	pb "code-assistant/proto/grpcbin"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("could not open stream: %v", err)
	}

	req := &pb.DummyMessage{
		FString: "test",
	}

	if err := stream.Send(req); err != nil {
		t.Fatalf("could not send message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("could not receive message: %v", err)
	}

	assert.Equal(t, req.FString, resp.FString, "Response should match request")
}
