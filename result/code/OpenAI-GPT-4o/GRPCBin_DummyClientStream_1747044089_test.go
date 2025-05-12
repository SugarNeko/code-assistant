package grpcbin_test

import (
	"context"
	"testing"
	"time"

	pb "code-assistant/proto/grpcbin"

	"google.golang.org/grpc"
)

func TestDummyClientStream(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	stream, err := client.DummyClientStream(ctx)
	if err != nil {
		t.Fatalf("could not open stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		dummyMessage := &pb.DummyMessage{
			FString:  "test",
			FInt32:   int32(i),
			FEnum:    pb.DummyMessage_ENUM_1,
			FSub:     &pb.DummyMessage_Sub{FString: "sub_test"},
			FBool:    true,
			FInt64:   int64(i),
			FBytes:   []byte("test"),
			FFloat:   float32(i),
		}
		if err := stream.Send(dummyMessage); err != nil {
			t.Fatalf("could not send: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("could not receive: %v", err)
	}

	// Validating the response
	expected := &pb.DummyMessage{
		FString:  "test",
		FInt32:   int32(9),
		FEnum:    pb.DummyMessage_ENUM_1,
		FSub:     &pb.DummyMessage_Sub{FString: "sub_test"},
		FBool:    true,
		FInt64:   int64(9),
		FBytes:   []byte("test"),
		FFloat:   float32(9),
	}

	if reply.FString != expected.FString || reply.FInt32 != expected.FInt32 {
		t.Errorf("unexpected response: got %+v, want %+v", reply, expected)
	}
}
