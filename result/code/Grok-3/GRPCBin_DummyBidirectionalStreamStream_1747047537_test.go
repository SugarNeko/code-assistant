package grpcbin

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr     = "grpcb.in:9000"
	connectTimeout = 15 * time.Second
)

func TestGRPCBinService_DummyBidirectionalStreamStream(t *testing.T) {
	// Set up gRPC connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := NewGRPCBinClient(conn)

	// Create bidirectional stream
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create bidirectional stream: %v", err)
	}

	// Test data to send
	testMessage := &DummyMessage{
		FString:  "test-string",
		FStrings: []string{"test1", "test2"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    DummyMessage_ENUM_1,
		FEnums:   []DummyMessage_Enum{DummyMessage_ENUM_0, DummyMessage_ENUM_1},
		FSub: &DummyMessage_Sub{
			FString: "sub-test",
		},
		FSubs:    []*DummyMessage_Sub{{FString: "sub1"}, {FString: "sub2"}},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   1234567890,
		FInt64S:  []int64{1, 2, 3},
		FBytes:   []byte("test-bytes"),
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	// Send test message to server
	if err := stream.Send(testMessage); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive response from server
	response, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	// Validate server response
	if response.FString != testMessage.FString {
		t.Errorf("Expected FString to be %q, got %q", testMessage.FString, response.FString)
	}
	if len(response.FStrings) != len(testMessage.FStrings) {
		t.Errorf("Expected FStrings length to be %d, got %d", len(testMessage.FStrings), len(response.FStrings))
	}
	if response.FInt32 != testMessage.FInt32 {
		t.Errorf("Expected FInt32 to be %d, got %d", testMessage.FInt32, response.FInt32)
	}
	if response.FEnum != testMessage.FEnum {
		t.Errorf("Expected FEnum to be %v, got %v", testMessage.FEnum, response.FEnum)
	}
	if response.FSub.FString != testMessage.FSub.FString {
		t.Errorf("Expected FSub.FString to be %q, got %q", testMessage.FSub.FString, response.FSub.FString)
	}
	if response.FBool != testMessage.FBool {
		t.Errorf("Expected FBool to be %v, got %v", testMessage.FBool, response.FBool)
	}
	if response.FInt64 != testMessage.FInt64 {
		t.Errorf("Expected FInt64 to be %d, got %d", testMessage.FInt64, response.FInt64)
	}
	if string(response.FBytes) != string(testMessage.FBytes) {
		t.Errorf("Expected FBytes to be %q, got %q", testMessage.FBytes, response.FBytes)
	}
	if response.FFloat != testMessage.FFloat {
		t.Errorf("Expected FFloat to be %f, got %f", testMessage.FFloat, response.FFloat)
	}

	// Close the stream
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to close stream: %v", err)
	}
}
