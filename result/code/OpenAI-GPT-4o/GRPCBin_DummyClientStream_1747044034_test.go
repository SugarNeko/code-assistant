package grpcbin_test

import (
	"context"
	"crypto/tls"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"code-assistant/proto/grpcbin"
)

func TestDummyClientStream(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithBlock(),
		grpc.WithTimeout(15*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyClientStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		msg := &grpcbin.DummyMessage{
			FString:   "Test",
			FInt32:    int32(i),
			FEnum:     grpcbin.DummyMessage_ENUM_1,
			FSub:      &grpcbin.DummyMessage_Sub{FString: "SubTest"},
			FBool:     true,
			FInt64:    int64(i),
			FBytes:    []byte("BytesTest"),
			FFloat:    float32(1.23),
		}

		if err := stream.Send(msg); err != nil {
			t.Fatalf("Failed to send message: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive reply: %v", err)
	}

	if reply.FString != "Test" {
		t.Errorf("Expected '%s', got '%s'", "Test", reply.FString)
	}
	if reply.FInt32 != 9 {
		t.Errorf("Expected '%d', got '%d'", 9, reply.FInt32)
	}
	if reply.FSub.FString != "SubTest" {
		t.Errorf("Expected '%s', got '%s'", "SubTest", reply.FSub.FString)
	}
	if reply.FBool != true {
		t.Errorf("Expected '%t', got '%t'", true, reply.FBool)
	}
}
