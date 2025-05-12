package grpcbin_test

import (
	"context"
	"testing"
	"time"

	"code-assistant/proto/grpcbin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcServerAddress = "grpcb.in:9000"
	connectTimeout    = 15 * time.Second
)

func TestGRPCBinService_DummyBidirectionalStreamStream(t *testing.T) {
	// Set up connection to the gRPC server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcbin.NewGRPCBinClient(conn)

	// Create a bidirectional stream
	stream, err := client.DummyBidirectionalStreamStream(context.Background())
	if err != nil {
		t.Fatalf("Failed to create bidirectional stream: %v", err)
	}

	// Prepare test data for sending
	testMessage := &grpcbin.DummyMessage{
		FString:  "test-string",
		FStrings: []string{"test1", "test2"},
		FInt32:   42,
		FInt32S:  []int32{1, 2, 3},
		FEnum:    grpcbin.DummyMessage_ENUM_1,
		FEnums:   []grpcbin.DummyMessage_Enum{grpcbin.DummyMessage_ENUM_0, grpcbin.DummyMessage_ENUM_2},
		FSub: &grpcbin.DummyMessage_Sub{
			FString: "sub-test-string",
		},
		FSubs: []*grpcbin.DummyMessage_Sub{
			{FString: "sub1"},
			{FString: "sub2"},
		},
		FBool:    true,
		FBools:   []bool{true, false},
		FInt64:   123456789,
		FInt64S:  []int64{1, 2, 3},
		FBytes:   []byte("test-bytes"),
		FBytess:  [][]byte{[]byte("bytes1"), []byte("bytes2")},
		FFloat:   3.14,
		FFloats:  []float32{1.1, 2.2, 3.3},
	}

	// Send the test message to the server
	if err := stream.Send(testMessage); err != nil {
		t.Fatalf("Failed to send message to server: %v", err)
	}

	// Receive and validate the response from the server
	resp, err := stream.Recv()
	if err != nil {
		t.Fatalf("Failed to receive response from server: %v", err)
	}

	// Validate the received response matches the sent data (echo behavior expected)
	if resp.FString != testMessage.FString {
		t.Errorf("Expected FString to be %q, got %q", testMessage.FString, resp.FString)
	}
	if len(resp.FStrings) != len(testMessage.FStrings) {
		t.Errorf("Expected FStrings length to be %d, got %d", len(testMessage.FStrings), len(resp.FStrings))
	}
	if resp.FInt32 != testMessage.FInt32 {
		t.Errorf("Expected FInt32 to be %d, got %d", testMessage.FInt32, resp.FInt32)
	}
	if resp.FEnum != testMessage.FEnum {
		t.Errorf("Expected FEnum to be %v, got %v", testMessage.FEnum, resp.FEnum)
	}
	if resp.FSub.FString != testMessage.FSub.FString {
		t.Errorf("Expected FSub.FString to be %q, got %q", testMessage.FSub.FString, resp.FSub.FString)
	}
	if resp.FBool != testMessage.FBool {
		t.Errorf("Expected FBool to be %v, got %v", testMessage.FBool, resp.FBool)
	}
	if resp.FInt64 != testMessage.FInt64 {
		t.Errorf("Expected FInt64 to be %d, got %d", testMessage.FInt64, resp.FInt64)
	}
	if string(resp.FBytes) != string(testMessage.FBytes) {
		t.Errorf("Expected FBytes to be %q, got %q", testMessage.FBytes, resp.FBytes)
	}
	if resp.FFloat != testMessage.FFloat {
		t.Errorf("Expected FFloat to be %f, got %f", testMessage.FFloat, resp.FFloat)
	}

	// Close the send direction of the stream
	if err := stream.CloseSend(); err != nil {
		t.Fatalf("Failed to close send stream: %v", err)
	}
}
