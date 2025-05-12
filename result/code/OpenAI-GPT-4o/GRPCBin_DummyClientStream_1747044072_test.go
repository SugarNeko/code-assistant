package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	// Create and send DummyMessages
	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:  "string",
			FInt32:   int32(i),
			FEnum:    grpcbin.DummyMessage_ENUM_1,
			FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
			FBool:    true,
			FInt64:   int64(i),
			FFloat:   1.0,
			FBytes:   []byte{0x01, 0x02},
		}
		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate client response
	expectedResponse := &grpcbin.DummyMessage{
		FString:  "string",
		FInt32:   9,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub"},
		FBool:    true,
		FInt64:   9,
		FFloat:   1.0,
		FBytes:   []byte{0x01, 0x02},
	}

	if reply.FString != expectedResponse.FString ||
		reply.FInt32 != expectedResponse.FInt32 ||
		reply.FEnum != expectedResponse.FEnum ||
		reply.FSub.FString != expectedResponse.FSub.FString ||
		reply.FBool != expectedResponse.FBool ||
		reply.FInt64 != expectedResponse.FInt64 ||
		reply.FFloat != expectedResponse.FFloat ||
		string(reply.FBytes) != string(expectedResponse.FBytes) {
		t.Errorf("Response does not match expected values")
	}
}
