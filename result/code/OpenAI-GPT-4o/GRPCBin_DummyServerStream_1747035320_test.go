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

func TestGRPCBinDummyServerStream(t *testing.T) {
	conn, err := grpc.Dial(
		"grpcb.in:9000",
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithBlock(),
		grpc.WithTimeout(15 * time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	req := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   42,
		FFloat:   3.14,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-test"},
		FInt64:   123456789,
		FBytes:   []byte("bytes"),
		FStrings: []string{"one", "two"},
	}

	stream, err := client.DummyServerStream(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call DummyServerStream: %v", err)
	}

	// Validate server responses
	for i := 0; i < 10; i++ {
		resp, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive message: %v", err)
		}

		// Validate response fields
		if resp.FString != req.FString {
			t.Errorf("Unexpected FString: got %v, want %v", resp.FString, req.FString)
		}
		if resp.FInt32 != req.FInt32 {
			t.Errorf("Unexpected FInt32: got %v, want %v", resp.FInt32, req.FInt32)
		}
		if resp.FFloat != req.FFloat {
			t.Errorf("Unexpected FFloat: got %v, want %v", resp.FFloat, req.FFloat)
		}
		if resp.FEnum != req.FEnum {
			t.Errorf("Unexpected FEnum: got %v, want %v", resp.FEnum, req.FEnum)
		}
		if resp.FSub.FString != req.FSub.FString {
			t.Errorf("Unexpected FSub.FString: got %v, want %v", resp.FSub.FString, req.FSub.FString)
		}
		if resp.FInt64 != req.FInt64 {
			t.Errorf("Unexpected FInt64: got %v, want %v", resp.FInt64, req.FInt64)
		}
		if string(resp.FBytes) != string(req.FBytes) {
			t.Errorf("Unexpected FBytes: got %v, want %v", resp.FBytes, req.FBytes)
		}
	}
}
