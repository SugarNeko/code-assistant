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
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("could not open stream: %v", err)
	}

	// Sending 10 sample messages
	for i := 0; i < 10; i++ {
		msg := &pb.DummyMessage{
			FString:   "test",
			FInt32:    int32(i),
			FEnum:     pb.DummyMessage_ENUM_1,
			FBool:     true,
			FInt64:    int64(i),
			FBytes:    []byte("sample bytes"),
			FFloat:    1.23,
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("could not send message: %v", err)
		}
	}

	// Closing the send direction
	res, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("could not receive response: %v", err)
	}

	// Validate response
	expected := &pb.DummyMessage{
		FString:   "test",
		FInt32:    int32(9),
		FEnum:     pb.DummyMessage_ENUM_1,
		FBool:     true,
		FInt64:    int64(9),
		FBytes:    []byte("sample bytes"),
		FFloat:    1.23,
	}

	if res.FString != expected.FString ||
		res.FInt32 != expected.FInt32 ||
		res.FEnum != expected.FEnum ||
		res.FBool != expected.FBool ||
		res.FInt64 != expected.FInt64 ||
		string(res.FBytes) != string(expected.FBytes) ||
		res.FFloat != expected.FFloat {
		t.Fatalf("unexpected response: got %v, expected %v", res, expected)
	}
}
