package grpcbintest

import (
	"context"
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
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create stream: %v", err)
	}

	testMessage := &grpcbin.DummyMessage{
		FString:  "test",
		FInt32:   123,
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FBool:    true,
		FInt64:   123456789,
		FFloat:   1.23,
		FSub:     &grpcbin.DummyMessage_Sub{FString: "sub-test"},
		FBytes:   []byte("testbytes"),
	}

	if err := stream.Send(testMessage); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	response, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	if response.FString != testMessage.FString {
		t.Errorf("Expected FString %s, got %s", testMessage.FString, response.FString)
	}
	if response.FInt32 != testMessage.FInt32 {
		t.Errorf("Expected FInt32 %d, got %d", testMessage.FInt32, response.FInt32)
	}
	if response.FEnum != testMessage.FEnum {
		t.Errorf("Expected FEnum %v, got %v", testMessage.FEnum, response.FEnum)
	}
	if response.FBool != testMessage.FBool {
		t.Errorf("Expected FBool %v, got %v", testMessage.FBool, response.FBool)
	}
	if response.FInt64 != testMessage.FInt64 {
		t.Errorf("Expected FInt64 %d, got %d", testMessage.FInt64, response.FInt64)
	}
	if response.FFloat != testMessage.FFloat {
		t.Errorf("Expected FFloat %f, got %f", testMessage.FFloat, response.FFloat)
	}
	if response.FSub.FString != testMessage.FSub.FString {
		t.Errorf("Expected FSub.FString %s, got %s", testMessage.FSub.FString, response.FSub.FString)
	}
	if string(response.FBytes) != string(testMessage.FBytes) {
		t.Errorf("Expected FBytes %s, got %s", string(testMessage.FBytes), string(response.FBytes))
	}
}
