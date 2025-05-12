package grpcbin_test

import (
	"context"
	"log"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
)

func TestDummyBidirectionalStreamStream(t *testing.T) {
	conn, err := grpc.Dial("grpcb.in:9000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(15*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to start bidirectional stream: %v", err)
	}

	dummyMessage := &grpcbin.DummyMessage{
		FString: "Test",
		FInt32:  123,
		FEnum:   grpcbin.DummyMessage_ENUM_1,
		FSub:    &grpcbin.DummyMessage_Sub{FString: "SubTest"},
	}

	if err := stream.Send(dummyMessage); err != nil {
		t.Fatalf("Failed to send dummy message: %v", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	if resp.FString != dummyMessage.FString {
		t.Errorf("Expected f_string %v, got %v", dummyMessage.FString, resp.FString)
	}
	if resp.FInt32 != dummyMessage.FInt32 {
		t.Errorf("Expected f_int32 %v, got %v", dummyMessage.FInt32, resp.FInt32)
	}
	if resp.FEnum != dummyMessage.FEnum {
		t.Errorf("Expected f_enum %v, got %v", dummyMessage.FEnum, resp.FEnum)
	}
	if resp.FSub.FString != dummyMessage.FSub.FString {
		t.Errorf("Expected f_sub.f_string %v, got %v", dummyMessage.FSub.FString, resp.FSub.FString)
	}
}
